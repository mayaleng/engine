package utils

import (
	"bytes"
	"log"
	"text/template"

	"mayaleng.org/engine/internal/platform/linguakit"
)

// ReplaceValues receive a Go template an replace values using the given list
// of words
func ReplaceValues(templateString string, words []linguakit.Word) string {
	input := map[string]interface{}{
		"Words": words,
	}

	var buf bytes.Buffer
	compiledTemplate := template.Must(template.New("tmp").Parse(templateString))

	error := compiledTemplate.Execute(&buf, input)

	if error != nil {
		log.Println(error)
		return templateString
	}

	return buf.String()
}
