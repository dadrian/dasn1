package parser

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (t *testing.T) TestParseSequence() {
	in := `Certificate  ::=  SEQUENCE  {
	tbsCertificate       TBSCertificate,
	signatureAlgorithm   AlgorithmIdentifier,
	signatureValue       BIT STRING  }`
	is := antlr.NewInputStream(in)
	lexer := NewASN1SchemaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewASN1SchemaParser(stream)
}
