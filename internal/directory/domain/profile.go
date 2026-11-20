package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
	_ "time/tzdata"
)

type ServiceType string

const (
	ServiceInPerson ServiceType = "in_person"
	ServiceRemote   ServiceType = "remote"
	ServiceHybrid   ServiceType = "hybrid"
)

func (s ServiceType) Valid() bool {
	return s == ServiceInPerson || s == ServiceRemote || s == ServiceHybrid
}

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
	StatusPaused    Status = "paused"
	StatusSuspended Status = "suspended"
)

type Verification string

const (
	VerificationUnverified       Verification = "unverified"
	VerificationContactVerified  Verification = "contact_verified"
	VerificationIdentityVerified Verification = "identity_verified"
)

type ContactPreference string

const (
	ContactPlatform ContactPreference = "platform"
	ContactEmail    ContactPreference = "email"
	ContactPhone    ContactPreference = "phone"
)

func (c ContactPreference) Valid() bool {
	return c == ContactPlatform || c == ContactEmail || c == ContactPhone
}

// Window is a weekly availability range. End must be after Start on the same
// day; overnight ranges are deliberately rejected in the MVP.
type Window struct {
	Day      int    `json:"day"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Timezone string `json:"timezone"`
}

type Profile struct {
	ID                        string
	AccountID                 string
	DisplayName               string
	Bio                       string
	ServiceType               ServiceType
	Status                    Status
	City                      string
	Neighborhood              string
	PublicAreaLabel           string
	RadiusKm                  *int
	PreciseLatEnc             string
	PreciseLngEnc             string
	PreciseLocationExpiresAt  *time.Time
	Availability              []Window
	AvailabilityNote          string
	ContactPreferenceOverride *ContactPreference
	Verification              Verification
	CategoryIDs               []string
	PublishedAt               *time.Time
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	SuspendedAt               *time.Time
	SuspendedReason           string
	Version                   int
}

func (p *Profile) HasPreciseLocation() bool {
	return p != nil && p.PreciseLatEnc != "" && p.PreciseLngEnc != ""
}

func (p *Profile) ValidateForPublish() error {
	if err := ValidateDisplayName(p.DisplayName); err != nil {
		return err
	}
	if len([]rune(p.Bio)) > 500 {
		return errors.New("a bio deve ter no máximo 500 caracteres")
	}
	if !p.ServiceType.Valid() {
		return errors.New("escolha um tipo de atendimento válido")
	}
	if len(p.CategoryIDs) == 0 {
		return errors.New("escolha ao menos uma categoria")
	}
	if err := ValidateServiceArea(p.ServiceType, p.City, p.Neighborhood, p.RadiusKm); err != nil {
		return err
	}
	if err := ValidateWindows(p.Availability); err != nil {
		return err
	}
	if p.ContactPreferenceOverride != nil && !p.ContactPreferenceOverride.Valid() {
		return errors.New("preferência de contato inválida")
	}
	return nil
}

func ValidateDisplayName(name string) error {
	name = strings.TrimSpace(name)
	if len([]rune(name)) < 2 || len([]rune(name)) > 80 {
		return errors.New("o nome de exibição deve ter 2 a 80 caracteres")
	}
	if strings.ContainsAny(name, "\r\n\t") {
		return errors.New("o nome de exibição contém caracteres inválidos")
	}
	return nil
}

func ValidateServiceArea(serviceType ServiceType, city, neighborhood string, radiusKm *int) error {
	if !serviceType.Valid() {
		return errors.New("tipo de atendimento inválido")
	}
	if serviceType == ServiceRemote {
		if radiusKm != nil && (*radiusKm < 1 || *radiusKm > 50) {
			return errors.New("o raio deve ficar entre 1 e 50 km")
		}
		return nil
	}
	if strings.TrimSpace(city) == "" || strings.TrimSpace(neighborhood) == "" {
		return errors.New("informe cidade e bairro para atendimento presencial")
	}
	if radiusKm == nil {
		return errors.New("informe o raio de atendimento")
	}
	if *radiusKm < 1 || *radiusKm > 50 {
		return errors.New("o raio deve ficar entre 1 e 50 km")
	}
	return nil
}

// ValidateWindows checks day/time/timezone validity and rejects overlapping
// windows. All windows must share one timezone so overlap semantics are clear.
func ValidateWindows(windows []Window) error {
	if len(windows) == 0 {
		return nil
	}
	if len(windows) > 21 {
		return errors.New("no máximo 21 janelas de disponibilidade")
	}
	var tz string
	for i, w := range windows {
		if w.Day < 0 || w.Day > 6 {
			return fmt.Errorf("janela %d: dia da semana inválido", i+1)
		}
		if _, err := time.Parse("15:04", w.Start); err != nil {
			return fmt.Errorf("janela %d: hora de início inválida", i+1)
		}
		if _, err := time.Parse("15:04", w.End); err != nil {
			return fmt.Errorf("janela %d: hora de fim inválida", i+1)
		}
		start, _ := time.Parse("15:04", w.Start)
		end, _ := time.Parse("15:04", w.End)
		if !end.After(start) {
			return fmt.Errorf("janela %d: o fim deve ser depois do início", i+1)
		}
		if strings.TrimSpace(w.Timezone) == "" {
			return fmt.Errorf("janela %d: informe o fuso horário", i+1)
		}
		if _, err := time.LoadLocation(w.Timezone); err != nil {
			return fmt.Errorf("janela %d: fuso horário inválido", i+1)
		}
		if tz == "" {
			tz = w.Timezone
		} else if tz != w.Timezone {
			return errors.New("todas as janelas devem usar o mesmo fuso horário")
		}
	}
	for i := 0; i < len(windows); i++ {
		for j := i + 1; j < len(windows); j++ {
			if windows[i].Day != windows[j].Day || windows[i].Timezone != windows[j].Timezone {
				continue
			}
			if overlap(windows[i], windows[j]) {
				return fmt.Errorf("as janelas %d e %d se sobrepõem", i+1, j+1)
			}
		}
	}
	return nil
}

func overlap(a, b Window) bool {
	as, _ := time.Parse("15:04", a.Start)
	ae, _ := time.Parse("15:04", a.End)
	bs, _ := time.Parse("15:04", b.Start)
	be, _ := time.Parse("15:04", b.End)
	return as.Before(be) && bs.Before(ae)
}

type ContactView struct {
	AccountPreference string
	ProfileOverride   string
	EffectiveChannel  string
	EmailMasked       string
	PhoneMasked       string
	PhoneVerified     bool
	PublicVisible     bool
}

// ResolveContact applies account -> profile precedence. An unverified phone is
// never an effective channel and never produces a public value.
func ResolveContact(accountPref, override, emailMasked, phoneMasked string, phoneVerified bool) ContactView {
	originalOverride := override
	effectivePref := override
	if effectivePref == "" {
		effectivePref = accountPref
	}
	pref := ContactPreference(effectivePref)
	if !pref.Valid() {
		pref = ContactPlatform
	}
	effective := pref
	if pref == ContactPhone && !phoneVerified {
		if emailMasked != "" && emailMasked != "***" {
			effective = ContactEmail
		} else {
			effective = ContactPlatform
		}
	}
	if pref == ContactEmail && (emailMasked == "" || emailMasked == "***") {
		effective = ContactPlatform
	}
	phone := phoneMasked
	if !phoneVerified {
		phone = ""
	}
	return ContactView{
		AccountPreference: accountPref,
		ProfileOverride:   originalOverride,
		EffectiveChannel:  string(effective),
		EmailMasked:       emailMasked,
		PhoneMasked:       phone,
		PhoneVerified:     phoneVerified,
		PublicVisible:     false,
	}
}

func ValidLatLng(lat, lng float64) bool {
	return lat >= -90 && lat <= 90 && lng >= -180 && lng <= 180 && !(lat == 0 && lng == 0)
}
