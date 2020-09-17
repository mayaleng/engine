package utils

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"mayaleng.org/engine/internal/platform/linguakit"
)

// ReplaceValues receive a Go template an replace values using the given list
// of words
func ReplaceValues(templateString string, words []linguakit.Word) string {
	var buf bytes.Buffer

	vars := map[string]linguakit.Word{}

	for i, word := range words {
		vars[fmt.Sprintf("Word%d", i+1)] = word
	}

	parsedTemplate, error := template.New("tmp").Funcs(template.FuncMap{
		"StartsWithVowel":     StartsWithVowel,
		"StartsWithConsonant": StartsWithConsonant,
		"GetFirstLetter":      GetFirstLetter,
		"ToLower":             ToLower,
		"GetKaqchikelNumber":  GetKaqchikelNumber,
	}).Option("missingkey=zero").Parse(templateString)

	if error != nil {
		return templateString
	}

	compiledTemplate := template.Must(parsedTemplate, error)

	error = compiledTemplate.Execute(&buf, vars)

	if error != nil {
		log.Println(error)
		return templateString
	}

	return buf.String()
}
