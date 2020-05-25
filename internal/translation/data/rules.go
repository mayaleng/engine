package data

// Detail represents a word in the rule
type Detail struct {
	Type       string
	Properties map[string]string
}

// Output represents a single rule to make an element of the
// translated sentence
type Output map[string]string

// TranslationRule represents a simple translation rule
type TranslationRule struct {
	Pattern string
	Details []Detail
	Output  []Output
}
