package schema

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"unicode"
)

type TokenType uint16

// TwoWordTokenFlag is a bit mask that indicates the TokenType is two words
const TwoWordTokenFlag TokenType = 0x8000

// RequiresPreviousTokenFlag is a bit mask that indicates the previous token is required for context
const RequiresPreviousTokenFlag TokenType = 0x4000

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
	DEFINED_BY_LITERAL        TokenType = 16 | TwoWordTokenFlag
	ANY_LITERAL               TokenType = 17
	OF_LITERAL                TokenType = 18 | RequiresPreviousTokenFlag
	CHOICE_LITERAL            TokenType = 19
	SIZE_LITERAL              TokenType = 20
	MAX                       TokenType = 21
	BIT_STRING_LITERAL        TokenType = 28 | TwoWordTokenFlag
	OCTET_STRING_LITERAL      TokenType = 29 | TwoWordTokenFlag
	OBJECT_IDENTIFIER_LITERAL TokenType = 30 | TwoWordTokenFlag
	DOTDOT                    TokenType = 32
	INTEGER                   TokenType = 33
	WORD                      TokenType = 34
)

type Token struct {
	Typ   TokenType
	Value string
}

//go:generate stringer -type=TokenType

type tokenState int

const (
	tokenStateBegin tokenState = iota
	tokenStateMaybeN
	tokenStateFinishLine
)

type runeAction int

const (
	runeActionIgnore runeAction = iota
	runeActionAppend
	runeActionSplitOne
)

func stateTransition(curState tokenState, r rune) (tokenState, runeAction) {
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
	case tokenStateFinishLine:
		if r == '\n' {
			return tokenStateBegin, runeActionIgnore
		}
		return tokenStateFinishLine, runeActionIgnore
	default:
		panic("invalid token state")
	}
}

func CollapseOrAppendTokens(tokens []Token, next *Token) (changed bool, replacement TokenType) {
	if next.Typ != WORD {
		return
	}
	tokensSoFar := len(tokens)
	if tokensSoFar < 1 {
		return
	}
	prev := &tokens[tokensSoFar-1]
	if next.Value == "STRING" && prev.Typ == WORD {
		switch prev.Value {
		case "BIT":
			return true, BIT_STRING_LITERAL
		case "OCTET":
			return true, OCTET_STRING_LITERAL
		}
		return
	}
	if next.Value == "IDENTIFIER" && prev.Typ == WORD {
		switch prev.Value {
		case "OBJECT":
			return true, OBJECT_IDENTIFIER_LITERAL
		}
		return
	}
	return
}

func ChunkToTokenType(chunk string) (ok bool, out TokenType) {
	switch chunk {
	case "::=":
		return true, ASSIGN
	case "{":
		return true, LBRACE
	case "}":
		return true, RBRACE
	case "[":
		return true, LBRACKET
	case "]":
		return true, RBRACKET
	case "(":
		return true, LPAREN
	case ")":
		return true, RPAREN
	case ",":
		return true, COMMA
	case "SEQUENCE":
		return true, SEQUENCE_LITERAL
	case "INTEGER":
		return true, INTEGER_LITERAL
	case "IMPLICIT":
		return true, IMPLICIT_LITERAL
	case "EXPLICIT":
		return true, EXPLICIT_LITERAL
	case "OPTIONAL":
		return true, OPTIONAL_LITERAL
	case "DEFAULT":
		return true, DEFAULT_LITERAL
	case "DEFINED BY":
		return true, DEFINED_BY_LITERAL
	case "ANY":
		return true, ANY_LITERAL
	case "OF":
		return true, OF_LITERAL
	case "CHOICE":
		return true, CHOICE_LITERAL
	case "SIZE":
		return true, SIZE_LITERAL
	case "MAX":
		return true, MAX
	case "BIT STRING":
		return true, BIT_STRING_LITERAL
	case "OCTET STRING":
		return true, OCTET_STRING_LITERAL
	case "OBJECT IDENTIFIER":
		return true, OBJECT_IDENTIFIER_LITERAL
	case "..":
		return true, DOTDOT
	}
	if strings.HasPrefix(chunk, "--") {
		return true, COMMENT
	}
	if strings.IndexFunc(chunk, func(r rune) bool {
		return !unicode.IsDigit(r)
	}) == -1 {
		return true, INTEGER
	}
	if strings.IndexFunc(chunk, func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r)
	}) == -1 {
		return true, WORD
	}
	return false, out
}

func finishChunk(chunk *[]rune, out *[]Token, nextRuneState tokenState) (nextState tokenState) {
	nextState = nextRuneState
	if len(*chunk) > 0 {
		c := string(*chunk)
		if ok, tt := ChunkToTokenType(c); ok {
			// TODO(dadrian): Only save values for token types that are variable
			if tt != COMMENT {
				t := Token{
					Typ:   tt,
					Value: c,
				}
				needsReplacement, newTokenType := CollapseOrAppendTokens(*out, &t)
				if needsReplacement {
					(*out)[len(*out)-1] = Token{
						Typ: newTokenType,
					}
				} else {
					*out = append(*out, t)
				}
			}
			nextState = tokenStateTransition(nextRuneState, tt)
		}
		*chunk = (*chunk)[:0]
	}
	return
}

func tokenStateTransition(nextRuneState tokenState, tokenType TokenType) tokenState {
	if tokenType == COMMENT {
		return tokenStateFinishLine
	}
	return nextRuneState
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
	for {
		r, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("not utf8")
		}
		nextState, action := stateTransition(state, r)
		switch action {
		case runeActionIgnore:
		case runeActionAppend:
			chunk = append(chunk, r)
		case runeActionSplitOne:
			// TODO(dadrian): I should really use a stack here
			nextState = finishChunk(&chunk, &out, nextState)
			chunk = append(chunk, r)
		}
		if nextState == tokenStateBegin {
			nextState = finishChunk(&chunk, &out, nextState)
		}
		state = nextState
	}
	return out, nil
}

type ASTNode struct {
	Token      *Token
	Parent     *ASTNode
	LeftChild  *ASTNode
	RightChild *ASTNode

	Body *ASTNode
}

type AST struct {
	Root *ASTNode
}

func rotateLeft(ast *AST, existingNode, newNode *ASTNode) {
	// TODO(dadrian): enforce well-formed constraints on tree?
	oldParent := existingNode.Parent
	newNode.Parent = oldParent
	if oldParent != nil {
		isLeft := existingNode == oldParent.LeftChild
		if isLeft {
			oldParent.LeftChild = newNode
		} else {
			oldParent.RightChild = newNode
		}
	} else {
		ast.Root = newNode
	}
	newNode.LeftChild = existingNode
	return
}

func consumeSequence(tokens []Token) (*ASTNode, error) {
	head := &tokens[0]
	switch head.Typ {
	case LBRACE:
	}
	return nil, nil
}

func addTokenToTree(ast *AST, lastNode *ASTNode, token *Token) (*ASTNode, error) {
	node := &ASTNode{
		Token: token,
	}
	switch token.Typ {
	case ASSIGN:
		rotateLeft(ast, lastNode, node)
		return node, nil
	case SEQUENCE_LITERAL:
		return node, nil
	}
	lastNode.RightChild = node
	node.Parent = lastNode
	return node, nil
}

func TokensToAST(tokens []Token) (*AST, error) {
	ast := AST{}
	if len(tokens) == 0 {
		return &ast, nil
	}
	ast.Root = &ASTNode{
		Token: &tokens[0],
	}
	lastNode := ast.Root
	for i := range tokens[1:] {
		node, err := addTokenToTree(&ast, lastNode, &tokens[i+1])
		if err != nil {
			return nil, err
		}
		lastNode = node
	}
	return &ast, nil
}

func (node *ASTNode) String() string {
	if node == nil {
		return "()"
	}
	return fmt.Sprintf("(%s %s %s)", node.Token.Typ, node.LeftChild.String(), node.RightChild.String())
}

func (ast *AST) String() string {
	return ast.Root.String()
}
