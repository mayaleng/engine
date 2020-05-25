package linguakit

// Word represent a word analyzed by Linguakit
//
// It is the representation of an analysis like this:
// perro_NOUN_1_<gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C|>
//
// Where Type=NOUN, Lemma=perro and Properties=gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C
//
type Word struct {
	Type       string            `json:"type"`
	Lemma      string            `json:"lemma"`
	Properties map[string]string `json:"properties"`
}
