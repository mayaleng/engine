package translator

import "mayaleng.org/engine/internal/platform/data"

// Translator represents the entity capable of translate words
type Translator struct {
	WordsHelper        data.WordsHelper
	TranslationsHelper data.TranslationsHelper
	RulesHelper        data.RulesHelper
}
