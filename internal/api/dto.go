package api

import (
	"time"

	catalogdomain "github.com/raire/raire/internal/catalog/domain"
	directoryapp "github.com/raire/raire/internal/directory/app"
	"github.com/raire/raire/internal/directory/domain"
	identitydomain "github.com/raire/raire/internal/identity/domain"
)

type accountResponse struct {
	ID                   string   `json:"id"`
	Email                string   `json:"email,omitempty"`
	DisplayName          string   `json:"display_name"`
	Role                 string   `json:"role"`
	Roles                []string `json:"roles"`
	EmailVerified        bool     `json:"email_verified"`
	ContactEmailMasked   string   `json:"contact_email_masked,omitempty"`
	ContactPhoneMasked   string   `json:"contact_phone_masked,omitempty"`
	ContactPhoneVerified bool     `json:"contact_phone_verified"`
	ContactPreference    string   `json:"contact_preference"`
}

func accountDTO(a *identitydomain.Account) accountResponse {
	roles := []string{string(a.Role)}
	if a.Role == identitydomain.RoleProfessional {
		roles = []string{string(identitydomain.RoleClient), string(identitydomain.RoleProfessional)}
	}
	resp := accountResponse{
		ID: a.ID, Email: a.EmailNormalized, DisplayName: a.DisplayName, Role: string(a.Role), Roles: roles,
		EmailVerified: a.EmailVerified(), ContactPhoneVerified: a.PhoneVerified(),
		ContactPreference: string(a.EffectiveContact()),
	}
	if a.ContactEmail != "" {
		resp.ContactEmailMasked = identitydomain.MaskEmail(a.ContactEmail)
	}
	if a.ContactPhone != "" {
		resp.ContactPhoneMasked = identitydomain.MaskPhone(a.ContactPhone)
	}
	return resp
}

type categoryDTO struct {
	ID          string   `json:"id"`
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Active      bool     `json:"active"`
	SortOrder   int      `json:"sort_order"`
	Synonyms    []string `json:"synonyms,omitempty"`
}

func categoryDTOFrom(c catalogdomain.Category) categoryDTO {
	return categoryDTO{
		ID: c.ID, Slug: c.Slug, Name: c.Name, Description: c.Description, Active: c.Active,
		SortOrder: c.SortOrder, Synonyms: c.Synonyms,
	}
}

type windowDTO struct {
	Day      int    `json:"day"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Timezone string `json:"timezone"`
}

func windowsFromDTO(in *[]windowDTO) []domain.Window {
	if in == nil {
		return nil
	}
	out := make([]domain.Window, 0, len(*in))
	for _, w := range *in {
		out = append(out, domain.Window{Day: w.Day, Start: w.Start, End: w.End, Timezone: w.Timezone})
	}
	return out
}

func windowsToDTO(in []domain.Window) []windowDTO {
	out := make([]windowDTO, 0, len(in))
	for _, w := range in {
		out = append(out, windowDTO{Day: w.Day, Start: w.Start, End: w.End, Timezone: w.Timezone})
	}
	return out
}

type categoryRefDTO struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type publicProfileResponse struct {
	ID               string           `json:"id"`
	DisplayName      string           `json:"display_name"`
	Bio              string           `json:"bio,omitempty"`
	ServiceType      string           `json:"service_type"`
	Categories       []categoryRefDTO `json:"categories"`
	City             string           `json:"city,omitempty"`
	Neighborhood     string           `json:"neighborhood,omitempty"`
	RadiusKm         *int             `json:"radius_km,omitempty"`
	Availability     []windowDTO      `json:"availability,omitempty"`
	AvailabilityNote string           `json:"availability_note,omitempty"`
	Verification     string           `json:"verification"`
	PublishedAt      *time.Time       `json:"published_at,omitempty"`
}

type contactResponse struct {
	AccountPreference string `json:"account_preference"`
	ProfileOverride   string `json:"profile_override"`
	EffectiveChannel  string `json:"effective_channel"`
	EmailMasked       string `json:"email_masked,omitempty"`
	PhoneMasked       string `json:"phone_masked,omitempty"`
	PhoneVerified     bool   `json:"phone_verified"`
	PublicVisible     bool   `json:"public_visible"`
}

type ownerProfileResponse struct {
	publicProfileResponse
	Status                   string          `json:"status"`
	CategoryIDs              []string        `json:"category_ids"`
	Contact                  contactResponse `json:"contact"`
	PreciseLocationSet       bool            `json:"precise_location_set"`
	PreciseLocationExpiresAt *time.Time      `json:"precise_location_expires_at,omitempty"`
	SuspendedReason          string          `json:"suspended_reason,omitempty"`
}

func publicViewDTO(v *directoryapp.PublicView) publicProfileResponse {
	categories := make([]categoryRefDTO, 0, len(v.Categories))
	for _, c := range v.Categories {
		categories = append(categories, categoryRefDTO{ID: c.ID, Slug: c.Slug, Name: c.Name})
	}
	return publicProfileResponse{
		ID: v.ID, DisplayName: v.DisplayName, Bio: v.Bio, ServiceType: v.ServiceType,
		Categories: categories, City: v.City, Neighborhood: v.Neighborhood, RadiusKm: v.RadiusKm,
		Availability: windowsToDTO(v.Availability), AvailabilityNote: v.AvailabilityNote,
		Verification: v.Verification, PublishedAt: v.PublishedAt,
	}
}

func ownerViewDTO(v *directoryapp.OwnerView) ownerProfileResponse {
	base := publicViewDTO(&v.PublicView)
	return ownerProfileResponse{
		publicProfileResponse: base, Status: v.Status, CategoryIDs: v.CategoryIDs,
		Contact: contactResponse{
			AccountPreference: v.Contact.AccountPreference, ProfileOverride: v.Contact.ProfileOverride,
			EffectiveChannel: v.Contact.EffectiveChannel, EmailMasked: v.Contact.EmailMasked,
			PhoneMasked: v.Contact.PhoneMasked, PhoneVerified: v.Contact.PhoneVerified,
			PublicVisible: v.Contact.PublicVisible,
		},
		PreciseLocationSet: v.PreciseLocationSet, PreciseLocationExpiresAt: v.PreciseLocationExpiresAt,
		SuspendedReason: v.SuspendedReason,
	}
}

type profileWriteRequest struct {
	DisplayName          *string      `json:"display_name"`
	Bio                  *string      `json:"bio"`
	ServiceType          *string      `json:"service_type"`
	CategoryIDs          *[]string    `json:"category_ids"`
	City                 *string      `json:"city"`
	Neighborhood         *string      `json:"neighborhood"`
	RadiusKm             *int         `json:"radius_km"`
	ClearRadius          bool         `json:"clear_radius"`
	Availability         *[]windowDTO `json:"availability"`
	AvailabilityNote     *string      `json:"availability_note"`
	ContactPreference    *string      `json:"contact_preference"`
	PreciseLat           *float64     `json:"precise_lat"`
	PreciseLng           *float64     `json:"precise_lng"`
	ClearPreciseLocation bool         `json:"clear_precise_location"`
}

type consentDTO struct {
	ID        string     `json:"id"`
	Kind      string     `json:"kind"`
	Version   string     `json:"version"`
	GrantedAt time.Time  `json:"granted_at"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

type adminProfileDTO struct {
	ID           string    `json:"id"`
	AccountID    string    `json:"account_id"`
	DisplayName  string    `json:"display_name"`
	Status       string    `json:"status"`
	City         string    `json:"city,omitempty"`
	Neighborhood string    `json:"neighborhood,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type auditDTO struct {
	ID            string         `json:"id"`
	ActorID       string         `json:"actor_id,omitempty"`
	ActorRole     string         `json:"actor_role,omitempty"`
	Action        string         `json:"action"`
	TargetType    string         `json:"target_type"`
	TargetID      string         `json:"target_id"`
	Reason        string         `json:"reason,omitempty"`
	CorrelationID string         `json:"correlation_id,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
}

type flagDTO struct {
	Key         string     `json:"key"`
	Enabled     bool       `json:"enabled"`
	Description string     `json:"description,omitempty"`
	Owner       string     `json:"owner,omitempty"`
	ReviewAt    *time.Time `json:"review_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
