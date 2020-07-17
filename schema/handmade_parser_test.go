package schema

import (
	"os"
	"testing"
)

func TestTokenizer(t *testing.T) {
	files := []string{
		"success/sequence.asn1",
		"success/sequence_sequence.asn1",
		"success/comments.asn1",
	}
	for _, p := range files {
		prefixedPath := "testdata/" + p
		f, err := os.Open(prefixedPath)
		if err != nil {
			t.Errorf("unable to open file: %s", prefixedPath)
			continue
		}
		tokens, err := Tokenize(f)
		t.Log(tokens)
		if err != nil {
			t.Errorf("failed to tokenize: %s", err)
			continue
		}
		ast, err := TokensToAST(tokens)
		t.Logf("%s", ast)
		if err != nil {
			t.Errorf("failed to make ast: %s", err)
			continue
		}
		t.Fail()
	}
}
