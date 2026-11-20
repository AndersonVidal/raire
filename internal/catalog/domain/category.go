package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type Category struct {
	ID          string
	Slug        string
	Name        string
	Description string
	Active      bool
	SortOrder   int
	Synonyms    []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var slugRE = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func NormalizeSlug(slug string) string {
	slug = strings.ToLower(strings.TrimSpace(slug))
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func ValidateSlug(slug string) error {
	if len(slug) < 2 || len(slug) > 60 || !slugRE.MatchString(slug) {
		return errors.New("slug deve ter 2 a 60 caracteres, letras minúsculas, números e hífen")
	}
	return nil
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if len([]rune(name)) < 2 || len([]rune(name)) > 80 {
		return errors.New("nome da categoria deve ter 2 a 80 caracteres")
	}
	if strings.ContainsAny(name, "\r\n\t") {
		return errors.New("nome da categoria contém caracteres inválidos")
	}
	return nil
}

func ValidateDescription(description string) error {
	if len([]rune(description)) > 500 {
		return errors.New("descrição deve ter no máximo 500 caracteres")
	}
	return nil
}

// normalizeText lowercases and removes common Portuguese accents so search is
// forgiving without a search engine.
func normalizeText(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ã", "a", "ä", "a",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"í", "i", "ì", "i", "î", "i", "ï", "i",
		"ó", "o", "ò", "o", "ô", "o", "õ", "o", "ö", "o",
		"ú", "u", "ù", "u", "û", "u", "ü", "u",
		"ç", "c",
	)
	s = replacer.Replace(s)
	return strings.Join(strings.Fields(s), " ")
}

func NormalizeTerm(term string) string { return normalizeText(term) }

func ValidateSynonyms(synonyms []string) error {
	if len(synonyms) > 20 {
		return errors.New("no máximo 20 sinônimos por categoria")
	}
	seen := map[string]bool{}
	for _, s := range synonyms {
		n := NormalizeTerm(s)
		if len([]rune(n)) < 2 || len([]rune(n)) > 60 {
			return errors.New("cada sinônimo deve ter 2 a 60 caracteres")
		}
		if seen[n] {
			return errors.New("sinônimos duplicados")
		}
		seen[n] = true
	}
	return nil
}

func SlugFromName(name string) string {
	n := normalizeText(name)
	var b strings.Builder
	lastDash := false
	for _, r := range n {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			lastDash = false
			continue
		}
		if !lastDash && b.Len() > 0 {
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}
