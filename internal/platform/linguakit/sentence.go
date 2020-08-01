package linguakit

// Sentence represents an array of words.
type Sentence struct {
	Pattern string
	Words   []Word
}

// Analyze represents a sentence to proccess by Linguakit
type Analyze struct {
	Phrase string
}
