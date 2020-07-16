package schema

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestTokenizer(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	files := []string{
		"success/sequence.asn1",
		"success/sequence_sequence.asn1",
	}
	for _, p := range files {
		prefixedPath := "testdata/" + p
		f, err := os.Open(prefixedPath)
		if err != nil {
			t.Errorf("unable to open file: %s", prefixedPath)
			continue
		}
		tokens, err := Tokenize(f)
		t.Fail()
		t.Log(tokens)
		t.Log(err)
	}
}
