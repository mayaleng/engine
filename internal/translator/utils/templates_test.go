package utils

import (
	"testing"

	"mayaleng.org/engine/internal/platform/linguakit"
)

func TestTemplates(t *testing.T) {
	t.Run("Replacing values with dot notation with valid input", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
			},
			{
				Lemma: "b",
			},
		}
		template := "My lemma is {{.Word1.Lemma}}"

		output := ReplaceValues(template, vars)

		if output != "My lemma is a" {
			t.Fatalf("Expecting %s, got %s", "My lemma is a", output)
		}
	})

	t.Run("Replacing with empty when input is incomplete", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
			},
			{
				Lemma: "b",
			},
		}
		template := "My lemma is {{.Word3.Lemma}}"

		output := ReplaceValues(template, vars)

		if output != "My lemma is " {
			t.Fatalf("Expecting %s, got %s", "My lemma is ", output)
		}
	})

	t.Run("Do nothing when an invalid template is given", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
			},
			{
				Lemma: "b",
			},
		}

		template := "{{unknownfunc}}"

		output := ReplaceValues(template, vars)

		if output != template {
			t.Fatalf("Expecting %s, got %s", template, output)
		}
	})

	t.Run("Replacing with the consequence when conditional operators are used", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
			},
			{
				Lemma: "b",
			},
		}

		template := "{{if and (eq .Word1.Lemma \"a\") (eq .Word2.Lemma \"b\") }} {{- .Word2.Lemma }} {{- else}} ok {{end}}"

		output := ReplaceValues(template, vars)

		if output != "b" {
			t.Fatalf("Expecting %s, got %s", "b", output)
		}
	})

	t.Run("Replacing with the alternative when conditional operators are used", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
			},
			{
				Lemma: "b",
			},
		}

		template := "{{if and (or (eq .Word2.Lemma \"aa\") (eq .Word2.Lemma \"a\")) (or (eq .Word2.Lemma \"bb\") (eq .Word2.Lemma \"b\")) }} {{- .Word2.Lemma -}} {{ else }} {{- \"z\" -}} {{ end}}"

		output := ReplaceValues(template, vars)

		if output != "z" {
			t.Fatalf("Expecting %s, got %s", "z", output)
		}
	})

	t.Run("Replacing nested in the consequence when conditional operators are used", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "a",
				Properties: map[string]string{
					"number": "S",
				},
			},
			{
				Lemma: "b",
			},
			{
				Lemma: "a",
				Properties: map[string]string{
					"number": "S",
				},
			},
		}

		template := "{{if and (eq .Word1.Properties.number \"S\") (eq .Word3.Properties.number \"S\")}} {{- \"a\" -}} {{else}} {{- \"r\" -}} {{end}}"

		output := ReplaceValues(template, vars)

		if output != "a" {
			t.Fatalf("Expecting %s, got %s", "a", output)
		}
	})
}