package app

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/raire/raire/internal/catalog/domain"
	"github.com/raire/raire/internal/contracts"
	"github.com/raire/raire/internal/platform/audit"
	"github.com/raire/raire/internal/platform/clock"
	"github.com/raire/raire/internal/platform/db"
	"github.com/raire/raire/internal/platform/id"
	"github.com/raire/raire/internal/platform/metrics"
	"github.com/raire/raire/internal/platform/outbox"
)

var ErrNotFound = errors.New("not found")
var ErrSlugTaken = errors.New("slug already taken")

type Repo interface {
	ListActive(ctx context.Context, q db.Querier) ([]domain.Category, error)
	ListAll(ctx context.Context, q db.Querier) ([]domain.Category, error)
	GetByID(ctx context.Context, q db.Querier, id string) (*domain.Category, error)
	GetBySlug(ctx context.Context, q db.Querier, slug string) (*domain.Category, error)
	Search(ctx context.Context, q db.Querier, normalizedTerm string) ([]domain.Category, error)
	Create(ctx context.Context, q db.Querier, c domain.Category) error
	Update(ctx context.Context, q db.Querier, c domain.Category) error
	ReplaceSynonyms(ctx context.Context, q db.Querier, categoryID string, synonyms []string, at time.Time) error
	Deactivate(ctx context.Context, q db.Querier, id string, at time.Time) error
}

type Actor struct {
	ID   string
	Role string
}

type CreateInput struct {
	Slug        string
	Name        string
	Description string
	Synonyms    []string
	SortOrder   int
}

type Service struct {
	uow     *db.UoW
	repo    Repo
	audit   *audit.Writer
	outbox  *outbox.Writer
	clock   clock.Clock
	metrics *metrics.Registry
}

func NewService(uow *db.UoW, repo Repo, auditWriter *audit.Writer, outboxWriter *outbox.Writer, c clock.Clock, m *metrics.Registry) *Service {
	if c == nil {
		c = clock.System{}
	}
	return &Service{uow: uow, repo: repo, audit: auditWriter, outbox: outboxWriter, clock: c, metrics: m}
}

func (s *Service) ListActive(ctx context.Context) ([]domain.Category, error) {
	return s.repo.ListActive(ctx, s.uow.Read())
}

func (s *Service) ListAll(ctx context.Context) ([]domain.Category, error) {
	return s.repo.ListAll(ctx, s.uow.Read())
}

func (s *Service) GetBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	c, err := s.repo.GetBySlug(ctx, s.uow.Read(), domain.NormalizeSlug(slug))
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	return c, nil
}

func (s *Service) Search(ctx context.Context, term string) ([]domain.Category, error) {
	term = domain.NormalizeTerm(term)
	if term == "" {
		return s.ListActive(ctx)
	}
	return s.repo.Search(ctx, s.uow.Read(), term)
}

func (s *Service) Create(ctx context.Context, actor Actor, in CreateInput, correlationID string) (*domain.Category, error) {
	slug := domain.NormalizeSlug(in.Slug)
	if slug == "" {
		slug = domain.SlugFromName(in.Name)
	}
	if err := domain.ValidateSlug(slug); err != nil {
		return nil, contracts.Validation(err.Error())
	}
	if err := domain.ValidateName(in.Name); err != nil {
		return nil, contracts.Validation(err.Error())
	}
	if err := domain.ValidateDescription(in.Description); err != nil {
		return nil, contracts.Validation(err.Error())
	}
	if err := domain.ValidateSynonyms(in.Synonyms); err != nil {
		return nil, contracts.Validation(err.Error())
	}
	if existing, err := s.repo.GetBySlug(ctx, s.uow.Read(), slug); err == nil && existing != nil {
		return nil, contracts.New(contracts.CodeConflict, "Já existe uma categoria com este slug.")
	}
	now := s.clock.Now()
	c := domain.Category{
		ID: id.New("cat"), Slug: slug, Name: strings.TrimSpace(in.Name),
		Description: strings.TrimSpace(in.Description), Active: true, SortOrder: in.SortOrder,
		Synonyms: cleanSynonyms(in.Synonyms), CreatedAt: now, UpdatedAt: now,
	}
	err := s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Create(ctx, q, c); err != nil {
			return err
		}
		if err := s.repo.ReplaceSynonyms(ctx, q, c.ID, c.Synonyms, now); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: actor.ID, ActorRole: actor.Role, Action: "catalog.category_created",
			TargetType: "category", TargetID: c.ID, CorrelationID: correlationID,
			Metadata: map[string]any{"slug": c.Slug}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "catalog.category_changed", map[string]any{
			"category_id": c.ID, "change": "created", "occurred_at": now,
		})
	})
	if err != nil {
		return nil, err
	}
	s.count("category_created")
	return &c, nil
}

type UpdateInput struct {
	Name        *string
	Description *string
	Synonyms    *[]string
	SortOrder   *int
}

func (s *Service) Update(ctx context.Context, actor Actor, categoryID string, in UpdateInput, correlationID string) (*domain.Category, error) {
	c, err := s.repo.GetByID(ctx, s.uow.Read(), categoryID)
	if err != nil {
		return nil, contracts.New(contracts.CodeNotFound, "")
	}
	if in.Name != nil {
		if err := domain.ValidateName(*in.Name); err != nil {
			return nil, contracts.Validation(err.Error())
		}
		c.Name = strings.TrimSpace(*in.Name)
	}
	if in.Description != nil {
		if err := domain.ValidateDescription(*in.Description); err != nil {
			return nil, contracts.Validation(err.Error())
		}
		c.Description = strings.TrimSpace(*in.Description)
	}
	if in.Synonyms != nil {
		if err := domain.ValidateSynonyms(*in.Synonyms); err != nil {
			return nil, contracts.Validation(err.Error())
		}
		c.Synonyms = cleanSynonyms(*in.Synonyms)
	}
	if in.SortOrder != nil {
		c.SortOrder = *in.SortOrder
	}
	now := s.clock.Now()
	c.UpdatedAt = now
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Update(ctx, q, *c); err != nil {
			return err
		}
		if in.Synonyms != nil {
			if err := s.repo.ReplaceSynonyms(ctx, q, c.ID, c.Synonyms, now); err != nil {
				return err
			}
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: actor.ID, ActorRole: actor.Role, Action: "catalog.category_updated",
			TargetType: "category", TargetID: c.ID, CorrelationID: correlationID,
			Metadata: map[string]any{"slug": c.Slug}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "catalog.category_changed", map[string]any{
			"category_id": c.ID, "change": "updated", "occurred_at": now,
		})
	})
	if err != nil {
		return nil, err
	}
	s.count("category_updated")
	return c, nil
}

func (s *Service) Deactivate(ctx context.Context, actor Actor, categoryID, correlationID string) error {
	c, err := s.repo.GetByID(ctx, s.uow.Read(), categoryID)
	if err != nil {
		return contracts.New(contracts.CodeNotFound, "")
	}
	if !c.Active {
		return nil
	}
	now := s.clock.Now()
	err = s.uow.Do(ctx, func(ctx context.Context, q db.Querier) error {
		if err := s.repo.Deactivate(ctx, q, c.ID, now); err != nil {
			return err
		}
		if err := s.audit.Write(ctx, q, audit.Event{
			ActorID: actor.ID, ActorRole: actor.Role, Action: "catalog.category_deactivated",
			TargetType: "category", TargetID: c.ID, CorrelationID: correlationID,
			Metadata: map[string]any{"slug": c.Slug}, CreatedAt: now,
		}); err != nil {
			return err
		}
		return s.outbox.Publish(ctx, q, "catalog.category_changed", map[string]any{
			"category_id": c.ID, "change": "deactivated", "occurred_at": now,
		})
	})
	if err == nil {
		s.count("category_deactivated")
	}
	return err
}

type CategoryInfo struct {
	ID     string
	Slug   string
	Name   string
	Active bool
}

func (s *Service) ActiveCategory(ctx context.Context, categoryID string) (CategoryInfo, error) {
	c, err := s.repo.GetByID(ctx, s.uow.Read(), categoryID)
	if err != nil || !c.Active {
		return CategoryInfo{}, ErrNotFound
	}
	return CategoryInfo{ID: c.ID, Slug: c.Slug, Name: c.Name, Active: c.Active}, nil
}

func (s *Service) ActiveCategories(ctx context.Context, ids []string) ([]CategoryInfo, error) {
	out := make([]CategoryInfo, 0, len(ids))
	for _, id := range ids {
		info, err := s.ActiveCategory(ctx, id)
		if err != nil {
			return nil, err
		}
		out = append(out, info)
	}
	return out, nil
}

func cleanSynonyms(in []string) []string {
	out := make([]string, 0, len(in))
	seen := map[string]bool{}
	for _, s := range in {
		s = strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
		n := domain.NormalizeTerm(s)
		if s == "" || seen[n] {
			continue
		}
		seen[n] = true
		out = append(out, s)
	}
	return out
}

func (s *Service) count(name string) {
	if s.metrics != nil {
		s.metrics.Inc(name)
	}
}
