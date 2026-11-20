package domain

import "testing"

func TestSlugAndNormalization(t *testing.T) {
	if got := SlugFromName("Instalação Elétrica"); got != "instalacao-eletrica" {
		t.Fatalf("slug = %q", got)
	}
	if got := NormalizeTerm("  Instalação   Elétrica "); got != "instalacao eletrica" {
		t.Fatalf("normalized = %q", got)
	}
	for _, slug := range []string{"eletricista", "professor-particular", "a1"} {
		if err := ValidateSlug(slug); err != nil {
			t.Fatalf("%s rejected: %v", slug, err)
		}
	}
	for _, slug := range []string{"", "A", "com espaço", "com_underscore", "-comeca"} {
		if err := ValidateSlug(slug); err == nil {
			t.Fatalf("%q accepted", slug)
		}
	}
	if err := ValidateSynonyms([]string{"tomada", "fiação"}); err != nil {
		t.Fatal(err)
	}
	if err := ValidateSynonyms([]string{"tomada", "TOMADA"}); err == nil {
		t.Fatal("duplicate synonyms accepted")
	}
}
