package schema

import (
	"io/ioutil"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dadrian/dasn1/schema/parser"
	"github.com/sirupsen/logrus"
)

type testParseSuccessfulListener struct {
	*parser.BaseASN1SchemaListener

	t    *testing.T
	name string
}

func newTestParseSuccessfulListener(t *testing.T, name string) *testParseSuccessfulListener {
	return &testParseSuccessfulListener{
		t:    t,
		name: name,
	}
}

func (s *testParseSuccessfulListener) VisitErrorNode(node antlr.ErrorNode) {
	s.t.Errorf("failed to parse %s: %s", s.name, node.GetText())
}

func TestParseSequence(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	testFiles, _ := ioutil.ReadDir("testdata")
	count := 0
	for _, info := range testFiles {
		if info.IsDir() {
			t.Logf("skipping directory %s", info.Name())
			continue
		}
		count++
		filepath := info.Name()
		t.Logf("parsing %s", filepath)
		b, _ := ioutil.ReadFile(filepath)
		is := antlr.NewInputStream(string(b))
		lexer := parser.NewASN1SchemaLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewASN1SchemaParser(stream)
		p.BuildParseTrees = true
		listener := newTestParseSuccessfulListener(t, filepath)
		antlr.ParseTreeWalkerDefault.Walk(listener, p.Module())
	}
	if count == 0 {
		t.Error("no tests read")
	}
}
