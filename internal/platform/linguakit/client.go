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
	}
	*/
	result := `SENT::grito_VERB_0_<gender:0|lemma:gritar|mode:I|number:S|person:1|pos:0|tense:P|token:grito|type:M|> <blank>_SENT_1_<lemma:\<blank>|pos:1|token:\<blank>|>
---
`
	sentences, error := Parse(result)

	if error != nil {
		return []Sentence{}, error
	}

	return sentences, nil
}
