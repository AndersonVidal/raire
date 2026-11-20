package domain

import "testing"

func TestValidateWindows(t *testing.T) {
	valid := []Window{
		{Day: 1, Start: "08:00", End: "12:00", Timezone: "America/Sao_Paulo"},
		{Day: 1, Start: "13:00", End: "18:00", Timezone: "America/Sao_Paulo"},
	}
	if err := ValidateWindows(valid); err != nil {
		t.Fatalf("valid windows rejected: %v", err)
	}
	overlap := []Window{
		{Day: 1, Start: "08:00", End: "12:00", Timezone: "America/Sao_Paulo"},
		{Day: 1, Start: "11:00", End: "13:00", Timezone: "America/Sao_Paulo"},
	}
	if err := ValidateWindows(overlap); err == nil {
		t.Fatal("overlap accepted")
	}
	if err := ValidateWindows([]Window{{Day: 1, Start: "18:00", End: "08:00", Timezone: "America/Sao_Paulo"}}); err == nil {
		t.Fatal("end before start accepted")
	}
	if err := ValidateWindows([]Window{{Day: 9, Start: "08:00", End: "09:00", Timezone: "America/Sao_Paulo"}}); err == nil {
		t.Fatal("invalid day accepted")
	}
	if err := ValidateWindows([]Window{{Day: 1, Start: "08:00", End: "09:00", Timezone: "Mars/Olympus"}}); err == nil {
		t.Fatal("invalid timezone accepted")
	}
	mixed := []Window{
		{Day: 1, Start: "08:00", End: "09:00", Timezone: "America/Sao_Paulo"},
		{Day: 2, Start: "08:00", End: "09:00", Timezone: "UTC"},
	}
	if err := ValidateWindows(mixed); err == nil {
		t.Fatal("mixed timezones accepted")
	}
}

func TestValidateServiceArea(t *testing.T) {
	ten := 10
	if err := ValidateServiceArea(ServiceRemote, "", "", nil); err != nil {
		t.Fatalf("remote without area rejected: %v", err)
	}
	if err := ValidateServiceArea(ServiceInPerson, "", "Centro", &ten); err == nil {
		t.Fatal("in-person without city accepted")
	}
	if err := ValidateServiceArea(ServiceHybrid, "Cidade", "", &ten); err == nil {
		t.Fatal("hybrid without neighborhood accepted")
	}
	if err := ValidateServiceArea(ServiceInPerson, "Cidade", "Centro", nil); err == nil {
		t.Fatal("in-person without radius accepted")
	}
	zero := 0
	if err := ValidateServiceArea(ServiceInPerson, "Cidade", "Centro", &zero); err == nil {
		t.Fatal("zero radius accepted")
	}
	fiftyOne := 51
	if err := ValidateServiceArea(ServiceInPerson, "Cidade", "Centro", &fiftyOne); err == nil {
		t.Fatal("radius above maximum accepted")
	}
}

func TestResolveContact(t *testing.T) {
	v := ResolveContact("email", "phone", "a***@example.test", "****8888", false)
	if v.EffectiveChannel != "email" {
		t.Fatalf("unverified phone became effective: %#v", v)
	}
	if v.PhoneMasked != "" {
		t.Fatal("unverified phone was exposed")
	}
	v = ResolveContact("email", "phone", "a***@example.test", "****8888", true)
	if v.EffectiveChannel != "phone" || v.PhoneMasked != "****8888" {
		t.Fatalf("verified phone not effective: %#v", v)
	}
	v = ResolveContact("platform", "", "a***@example.test", "", false)
	if v.EffectiveChannel != "platform" || v.PublicVisible {
		t.Fatalf("platform default wrong: %#v", v)
	}
}

func TestProfileValidateForPublish(t *testing.T) {
	ten := 10
	p := Profile{
		DisplayName: "Ana Souza", ServiceType: ServiceInPerson,
		City: "Cidade Piloto", Neighborhood: "Centro", RadiusKm: &ten,
		CategoryIDs: []string{"cat_eletricista"},
	}
	if err := p.ValidateForPublish(); err != nil {
		t.Fatalf("valid profile rejected: %v", err)
	}
	p.CategoryIDs = nil
	if err := p.ValidateForPublish(); err == nil {
		t.Fatal("profile without category accepted")
	}
}
