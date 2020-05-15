package util

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func PrettyParseTree(t antlr.Tree, p antlr.Parser, tokens antlr.TokenStream) string {
	return prettyTree(t, p, tokens, 0)
	//return antlr.TreesStringTree(t, p.GetRuleNames(), p)
}

func indentBuilder(b *strings.Builder, level int) {
	indent := "     "
	for i := 0; i < level; i++ {
		b.WriteString(indent)
	}
}

func prettyTree(t antlr.Tree, p antlr.Parser, tokens antlr.TokenStream, level int) string {
	sb := strings.Builder{}
	indentBuilder(&sb, level)
	s := antlr.TreesGetNodeText(t, p.GetRuleNames(), p)
	if t.GetChildCount() == 0 {
		sb.WriteString(":")
	}
	sb.WriteString(s)
	for i := 0; i < t.GetChildCount(); i++ {
		sb.WriteRune('\n')
		sb.WriteString(prettyTree(t.GetChild(i), p, tokens, level+1))

	}
	indentBuilder(&sb, level)
	return sb.String()
}
