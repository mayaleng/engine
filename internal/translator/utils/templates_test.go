package utils

import (
	"testing"

	"mayaleng.org/engine/internal/platform/linguakit"
)

func TestTemplates(t *testing.T) {
	t.Run("X", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Properties: map[string]string{
					"tr":   "true",
					"intr": "true",
				},
			},
		}

		template := "{{ if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}w{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}nu{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}qa{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}aw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}iw{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}i{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}r{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}u{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}k{{ else if and ( .Word1.Properties.tr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}ki{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}in{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}i{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}oj{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}at{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}a{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"2\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}ix{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma )}}Ø{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithVowel .Word1.Lemma )}}e'{{ else if and ( .Word1.Properties.intr ) (eq .Word1.Properties.person \"3\") (eq .Word1.Properties.number \"P\") ( .Word1.StartWithConsonant .Word1.Lemma )}}e{{end}}"

		output := ReplaceValues(template, vars)

		if output != "tr-intr" {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

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

	t.Run("Validate if word starts with vowel with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "abrir",
			},
		}

		template := "{{if ( .Word1.StartWithVowel .Word1.Lemma )}} SI {{else}} NO {{end}}"

		output := ReplaceValues(template, vars)

		if output != " SI " {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word doesn't start with vowel with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "carro",
			},
		}

		template := "{{if not ( .Word1.StartWithVowel .Word1.Lemma )}} NO {{else}} SI {{end}}"

		output := ReplaceValues(template, vars)

		if output != " NO " {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word starts with vowel, person 1 and number S with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "arbol",
				Properties: map[string]string{
					"number": "S",
					"person": "1",
				},
			},
		}

		template := "{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithVowel .Word1.Lemma ) }}nu{{end}}"

		output := ReplaceValues(template, vars)

		if output != "nu" {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word starts with consonant, person 1 and number S with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "perro",
				Properties: map[string]string{
					"number": "S",
					"person": "1",
				},
			},
		}

		template := "{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( .Word1.StartWithConsonant .Word1.Lemma ) }}u{{end}}"

		output := ReplaceValues(template, vars)

		if output != "u" {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word is transitive and intransitive with success when using properties", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Properties: map[string]string{
					"tr":   "true",
					"intr": "true",
				},
			},
		}

		template := "{{ if and ( .Word1.Properties.tr ) ( .Word1.Properties.intr ) }}tr-intr{{end}}"

		output := ReplaceValues(template, vars)

		if output != "tr-intr" {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})
}
