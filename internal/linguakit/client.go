package linguakit

import (
	"bytes"
	"fmt"
	"os/exec"
)

func anaylizeSentence(setence string) (string, error) {
	var outbuf, errbuf bytes.Buffer

	cmd := exec.Command("linguakit", "dep", "es", setence, "-s")
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	errorRunning := cmd.Run()
	result := outbuf.String()
	error := errbuf.String()

	if errorRunning != nil {
		return "", fmt.Errorf("%s", error)
	}

	return result, nil
}
