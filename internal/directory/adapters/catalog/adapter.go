package catalog

import (
	"context"

	catalogapp "github.com/raire/raire/internal/catalog/app"
	directoryapp "github.com/raire/raire/internal/directory/app"
)

type Adapter struct {
	Service *catalogapp.Service
}

func New(service *catalogapp.Service) *Adapter { return &Adapter{Service: service} }

func (a *Adapter) ActiveCategory(ctx context.Context, categoryID string) (directoryapp.CategoryInfo, error) {
	info, err := a.Service.ActiveCategory(ctx, categoryID)
	if err != nil {
		return directoryapp.CategoryInfo{}, err
	}
	return directoryapp.CategoryInfo{ID: info.ID, Slug: info.Slug, Name: info.Name, Active: info.Active}, nil
}

func (a *Adapter) ActiveCategories(ctx context.Context, ids []string) ([]directoryapp.CategoryInfo, error) {
	infos, err := a.Service.ActiveCategories(ctx, ids)
	if err != nil {
		return nil, err
	}
	out := make([]directoryapp.CategoryInfo, 0, len(infos))
	for _, info := range infos {
		out = append(out, directoryapp.CategoryInfo{ID: info.ID, Slug: info.Slug, Name: info.Name, Active: info.Active})
	}
	return out, nil
}

var _ directoryapp.CatalogReader = (*Adapter)(nil)
