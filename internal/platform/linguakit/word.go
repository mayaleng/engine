package linguakit

import (
	"regexp"
	"strings"
)

// Word represent a word analyzed by Linguakit
//
// It is the representation of an analysis like this:
// perro_NOUN_1_<gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C|>
//
// Where Tag=NOUN, Lemma=perro and Properties=gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C
//
type Word struct {
	Tag         string            `json:"tag"`
	Lemma       string            `json:"lemma"`
	Translation string            `json:"translation"`
	Type        string            `json:"type"`
	Properties  map[string]string `json:"properties"`
}

// StartWithVowel validates if a word starts with vowel
func (w Word) StartWithVowel() bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return validPattern.MatchString(w.Translation)
}

// StartWithConsonant validates if a word starts with consonant
func (w Word) StartWithConsonant() bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return !validPattern.MatchString(w.Translation)
}

// FirstLetter returns the character number one of a string
func (w Word) FirstLetter() string {
	return w.Translation[0:1]
}

// ToLower returns the given string to lowercase
func (w Word) ToLower() string {
	return strings.ToLower(w.Lemma)
}
