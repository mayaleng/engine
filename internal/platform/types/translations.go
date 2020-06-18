package types

// TranslatedPhrase represents the translation of a phrase
type TranslatedPhrase struct {
	Phrase       string        `json:"phrase"`
	UnknownWords []UnknownWord `json:"unknown_words"`
}

// TranslatedSentence represents the translation of a single sentence
type TranslatedSentence struct {
	Sentence     string        `json:"sentence"`
	UnknownWords []UnknownWord `json:"unknown_words,omitempty"`
}

// UnknownWord is used as feedback in trasnlations
type UnknownWord struct {
	SourceLanguage string `json:"source,omitempty"`
	TargetLanguage string `json:"target,omitempty"`
	Word           string `json:"word,omitempty"`
}
