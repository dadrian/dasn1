package schema

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
)

type TokenType uint16

// Tokens used by ASN1 grammars
const (
	SPACE                     TokenType = 0
	COMMENT                   TokenType = 1
	ASSIGN                    TokenType = 2
	LBRACE                    TokenType = 3
	RBRACE                    TokenType = 4
	LBRACKET                  TokenType = 5
	RBRACKET                  TokenType = 6
	LPAREN                    TokenType = 7
	RPAREN                    TokenType = 8
	COMMA                     TokenType = 9
	SEQUENCE_LITERAL          TokenType = 10
	INTEGER_LITERAL           TokenType = 11
	IMPLICIT_LITERAL          TokenType = 12
	EXPLICIT_LITERAL          TokenType = 13
	OPTIONAL_LITERAL          TokenType = 14
	DEFAULT_LITERAL           TokenType = 15
	DEFINED_BY_LITERAL        TokenType = 16
	ANY_LITERAL               TokenType = 17
	OF_LITERAL                TokenType = 18
	CHOICE_LITERAL            TokenType = 19
	SIZE_LITERAL              TokenType = 20
	MAX                       TokenType = 21
	BMPSTRING_LITERAL         TokenType = 22
	IA5STRING_LITERAL         TokenType = 23
	PRINTABLE_STRING_LITERAL  TokenType = 24
	TELETEX_STRING_LITERAL    TokenType = 25
	UNIVERSAL_STRING_LITERAL  TokenType = 26
	UTF8_STRING_LITERAL       TokenType = 27
	BIT_STRING_LITERAL        TokenType = 28
	OCTET_STRING_LITERAL      TokenType = 29
	OBJECT_IDENTIFIER_LITERAL TokenType = 30
	DOTDOT                    TokenType = 31
	INTEGER                   TokenType = 32
	WORD                      TokenType = 33
)

type Token struct {
	Typ   TokenType
	Value string
}

//go:generate stringer -type=TokenType

type tokenState int

const (
	tokenStateBegin  tokenState = iota
	tokenStateMaybeN tokenState = iota
)

type runeAction int

const (
	runeActionIgnore   runeAction = iota
	runeActionAppend   runeAction = iota
	runeActionSplitOne runeAction = iota
)

func transition(curState tokenState, r rune) (tokenState, runeAction) {
	switch curState {
	case tokenStateBegin:
		if unicode.IsSpace(r) {
			return tokenStateBegin, runeActionIgnore
		}
		return tokenStateMaybeN, runeActionAppend
	case tokenStateMaybeN:
		if unicode.IsSpace(r) {
			return tokenStateBegin, runeActionIgnore
		}
		if r == ',' {
			return tokenStateBegin, runeActionSplitOne
		}
		return tokenStateMaybeN, runeActionAppend
	default:
		panic("invalid token state")
	}
}

func ChunkToTokenType(chunk string) (ok bool, out TokenType) {
	switch chunk {
	case "::=":
		return true, ASSIGN
	case "{":
		return true, LBRACE
	case "}":
		return true, RBRACE
	case "(":
		return true, LPAREN
	case ")":
		return true, RPAREN
	case ",":
		return true, COMMA
	case "SEQUENCE":
		return true, SEQUENCE_LITERAL
	}
	if strings.IndexFunc(chunk, func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r)
	}) == -1 {
		return true, WORD
	}
	return false, out
}

var chunks []string

func finishChunk(chunk *[]rune, out *[]Token) {
	if len(*chunk) > 0 {
		c := string(*chunk)
		logrus.Debugf("got chunk: %s", c)
		chunks = append(chunks, c)
		if ok, t := ChunkToTokenType(c); ok {
			// TODO(dadrian): Only save values for token types that are variable
			logrus.Debugf("got TokenType: %s", t)
			*out = append(*out, Token{
				Typ:   t,
				Value: c,
			})
		}
		*chunk = (*chunk)[:0]
	}
}

// Tokenize splits an input file into Tokens
func Tokenize(in io.ReadSeeker) ([]Token, error) {
	fileSize, err := in.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	if fileSize > math.MaxInt32 {
		return nil, errors.New("file too big")
	}
	out := make([]Token, 0, fileSize)
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	br := bufio.NewReaderSize(in, int(fileSize))
	state := tokenStateBegin
	chunk := make([]rune, 0)
	chunks = make([]string, 0)
	for {
		r, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("not utf8")
		}
		nextState, action := transition(state, r)
		switch action {
		case runeActionIgnore:
		case runeActionAppend:
			chunk = append(chunk, r)
		case runeActionSplitOne:
			// TODO(dadrian): I should really use a stack here
			finishChunk(&chunk, &out)
			chunk = append(chunk, r)
		}
		if nextState == tokenStateBegin {
			finishChunk(&chunk, &out)
		}
		state = nextState
	}
	logrus.Debugf("chunks [%d]: %v", len(chunks), chunks)
	return out, errors.New("unimplemented")
}
