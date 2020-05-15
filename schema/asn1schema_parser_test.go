package schema

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dadrian/dasn1/schema/parser"
	"github.com/dadrian/dasn1/util"
)

type testParseSuccessfulListener struct {
	antlr.DefaultErrorListener

	t    *testing.T
	name string

	failed bool
}

func (l *testParseSuccessfulListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.t.Errorf("%s: line %d:%d %s", l.name, line, column, msg)
	l.failed = true
}

func newTestParseSuccessfulListener(t *testing.T, name string) *testParseSuccessfulListener {
	return &testParseSuccessfulListener{
		t:    t,
		name: name,
	}
}

func testParseFile(t *testing.T, filepath string) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("could not open file %s", filepath)
	}
	is := antlr.NewInputStream(string(b))
	lexer := parser.NewASN1SchemaLexer(is)
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewASN1SchemaParser(stream)
	l := newTestParseSuccessfulListener(t, filepath)
	p.RemoveErrorListeners()
	p.AddErrorListener(l)
	p.BuildParseTrees = true
	tree := p.Module()
	s := util.PrettyParseTree(tree, tree.GetParser(), tree.GetParser().GetTokenStream())
	if l.failed {
		t.Logf("\n%s", s)
	}

}

func TestParseSequence(t *testing.T) {
	testFiles, _ := ioutil.ReadDir(path.Join("testdata", "success"))
	count := 0
	for _, info := range testFiles {
		if info.IsDir() {
			t.Logf("skipping directory %s", info.Name())
			continue
		}
		count++
		name := info.Name()
		t.Logf("parsing %s", name)
		filepath := path.Join("testdata", "success", name)
		testParseFile(t, filepath)
	}
	if count == 0 {
		t.Error("no tests read")
	}
}

func TestParseRFC5280(t *testing.T) {
	testParseFile(t, path.Join("testdata", "rfc", "rfc5280.asn1"))
}
