package linguakit

import (
	"regexp"
)

// Word represent a word analyzed by Linguakit
//
// It is the representation of an analysis like this:
// perro_NOUN_1_<gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C|>
//
// Where Tag=NOUN, Lemma=perro and Properties=gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C
//
type Word struct {
	Tag          string            `json:"tag"`
	Lemma        string            `json:"lemma"`
	Type         string            `json:"type"`
	Properties   map[string]string `json:"properties"`
	Transitive   bool              `json:"transitive"`
	Intransitive bool              `json:"intransitive"`
}

// StartWithVowel validates if a word starts with vowel
func (w Word) StartWithVowel(value string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return validPattern.MatchString(value)
}

// StartWithConsonant validates if a word starts with consonant
func (w Word) StartWithConsonant(value string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return !validPattern.MatchString(value)
}
