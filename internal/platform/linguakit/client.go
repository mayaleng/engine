package linguakit

// AnalyzePhrase get the given phrase and decompose it in an array
// of sentences
func AnalyzePhrase(phrase string) ([]Sentence, error) {
	/*var outbuf, errbuf bytes.Buffer

	cmd := exec.Command("linguakit", "dep", "es", phrase, "-s")
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	errorRunning := cmd.Run()
	result := outbuf.String()
	errorString := errbuf.String()

	if errorRunning != nil {
		return []Sentence{}, fmt.Errorf("%s", errorString)
	}*/

	result := `SENT::el_DET_0_<gender:M|lemma:el|number:S|person:0|pos:0|possessor:0|token:el|type:A|> unknown_NOUN_1_<gender:0|lemma:unknown|number:0|person:3|pos:1|token:unknown|type:C|> ._SENT_2_<lemma:.|pos:2|token:.|>
(SpecL;unknown_NOUN_1;el_DET_0)
---
`

	sentences, error := Parse(result)

	if error != nil {
		return []Sentence{}, error
	}

	return sentences, nil
}
