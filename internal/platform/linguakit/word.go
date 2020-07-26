package linguakit

/*
Word represents a word analyzed by Linguakit.
It's the representation of the Linguakit's result

e.g

Linguakit output:

	perro_NOUN_1_\<gender:M|lemma:perro|number:S|person:3|pos:1|token:perro|type:C|\>

Will become to:

	{ Tag: "Noun", Lemma: "perro", Type: "NOUN", Properties: map[string]string{"gender": "M"}}
*/
type Word struct {
	Tag        string            `json:"tag"`
	Lemma      string            `json:"lemma"`
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
}
