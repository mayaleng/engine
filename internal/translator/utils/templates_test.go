package utils

import (
	"testing"

	"mayaleng.org/engine/internal/platform/linguakit"
)

func TestTemplates(t *testing.T) {
	t.Run("Get kaqchikel number with success when using a cardinal number ", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma: "120",
			},
		}

		template := "{{ GetKaqchikelNumber .Word1.Lemma \"jun,ka'i',oxi',kaji',wo'o,waqi',wuqu',waqxaqi',b'eleje'\" \"ju,ka,ox,kaj,o,waq,wuq,waqxaq,b'elej,laj,julaj,kab'laj,oxlaj,kajlaj,wolaj,waqlaj,wuqlaj,waqxaqlaj,b'elejlaj\" \"k'al,q'o',chuy,k'ala'\" }}"

		output := ReplaceValues(template, vars)

		if output != "waqk'al" {
			t.Fatalf("The expected value for 120 would be waqk'al")
		}
	})

	t.Run("Applying ToLower and GetFirstLetter with success when using multiple words to evaluate", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Lemma:       "MeDio",
				Translation: "t'zi",
				Properties: map[string]string{
					"tr":     "true",
					"intr":   "true",
					"person": "1",
					"number": "S",
				},
			},
			{
				Translation: "vaca",
				Properties: map[string]string{
					"person": "1",
					"number": "S",
					"tr":     "true",
				},
			},
			{
				Lemma:       "regado",
				Translation: "regado",
				Properties: map[string]string{
					"tr": "true",
				},
			},
		}

		template := "{{ if (eq ( ToLower .Word1.Lemma ) \"muy\") }}{{.Word2.Translation}} {{.Word2.Translation}}{{ else if (eq ( ToLower .Word1.Lemma ) \"medio\") }}{{ .Word2.Translation }}{{ GetFirstLetter .Word2.Translation }}oj{{end}}"

		output := ReplaceValues(template, vars)

		if output != "vacavoj" {
			t.Fatalf("The expected value was vacavoj, it was retured %s", output)
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
				Translation: "abrir",
			},
		}

		template := "{{if ( StartsWithVowel .Word1.Translation )}} SI {{else}} NO {{end}}"

		output := ReplaceValues(template, vars)

		if output != " SI " {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word doesn't start with vowel with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Translation: "carro",
			},
		}

		template := "{{if not ( StartsWithVowel .Word1.Translation )}} NO {{else}} SI {{end}}"

		output := ReplaceValues(template, vars)

		if output != " NO " {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word starts with vowel, person 1 and number S with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Translation: "arbol",
				Properties: map[string]string{
					"number": "S",
					"person": "1",
				},
			},
		}

		template := "{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithVowel .Word1.Translation ) }}nu{{end}}"

		output := ReplaceValues(template, vars)

		if output != "nu" {
			t.Fatalf("Expecting %s, got %s", "SI", template)
		}
	})

	t.Run("Validate if word starts with consonant, person 1 and number S with success when using functions inside template", func(t *testing.T) {
		vars := []linguakit.Word{
			{
				Translation: "perro",
				Properties: map[string]string{
					"number": "S",
					"person": "1",
				},
			},
		}

		template := "{{ if and (eq .Word1.Properties.person \"1\") (eq .Word1.Properties.number \"S\") ( StartsWithConsonant .Word1.Translation ) }}u{{end}}"

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
