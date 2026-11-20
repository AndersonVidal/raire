package identity

import (
	"context"

	directoryapp "github.com/raire/raire/internal/directory/app"
	identityapp "github.com/raire/raire/internal/identity/app"
	"github.com/raire/raire/internal/platform/db"
)

// Adapter exposes the narrow identity read/write surface the directory needs.
type Adapter struct {
	Service *identityapp.Service
}

func New(service *identityapp.Service) *Adapter { return &Adapter{Service: service} }

func (a *Adapter) AccountInfo(ctx context.Context, accountID string) (directoryapp.AccountInfo, error) {
	info, err := a.Service.AccountInfo(ctx, accountID)
	if err != nil {
		return directoryapp.AccountInfo{}, err
	}
	return directoryapp.AccountInfo{
		ID: info.ID, DisplayName: info.DisplayName, EmailVerified: info.EmailVerified,
		ContactPreference: info.ContactPreference, ContactEmailMasked: info.ContactEmailMasked,
		ContactPhoneMasked: info.ContactPhoneMasked, PhoneVerified: info.PhoneVerified,
	}, nil
}

func (a *Adapter) MarkProfessional(ctx context.Context, accountID string) error {
	return a.Service.MarkProfessional(ctx, accountID)
}

func (a *Adapter) RevokeSessions(ctx context.Context, accountID string) error {
	return a.Service.RevokeSessions(ctx, accountID)
}

func (a *Adapter) RevokeSessionsTx(ctx context.Context, q db.Querier, accountID string) error {
	return a.Service.RevokeSessionsTx(ctx, q, accountID)
}

var _ directoryapp.AccountReader = (*Adapter)(nil)
