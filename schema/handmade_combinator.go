package schema

type TokenStream []Token

func (t TokenStream) head() *Token {
	return &t[0]
}

func (t TokenStream) Current() Token {
	return t[0]
}

func (t TokenStream) RemainingInput() TokenStream {
	return t[1:]
}

type ParserResult struct {
	Result         *ASTNode
	RemainingInput TokenStream
}

type Parser func(TokenStream) ParserResult

func ExpectToken(t TokenType) Parser {
	return func(input TokenStream) ParserResult {
		if input.Current().Typ == t {
			return ParserResult{
				Result: &ASTNode{
					Token: input.head(),
				},
				RemainingInput: input.RemainingInput(),
			}
		}
		return ParserResult{
			Result:         nil,
			RemainingInput: input,
		}
	}
}

func ReadUntil(t TokenType) Parser {
	return func(input TokenStream) ParserResult {
		it := input
		for {
			if len(it) == 0 {
				break
			}
			if it.Current().Typ == t {
				return ParserResult{
					Result: &ASTNode{
						Token: it.head(),
					},
					RemainingInput: it.RemainingInput(),
				}
			}
			it = it.RemainingInput()
		}
		return ParserResult{
			Result:         nil,
			RemainingInput: input,
		}
	}
}

func ExpectNotToken(ts []TokenType) Parser {
	return func(input TokenStream) ParserResult {
		for _, blockedToken := range ts {
			if input.head().Typ == blockedToken {
				return ParserResult{
					Result:         nil,
					RemainingInput: input,
				}
			}
		}
		return ParserResult{
			Result: &ASTNode{
				Token: input.head(),
			},
			RemainingInput: input.RemainingInput(),
		}
	}
}

//func ExpectTokensInOrder(ts []TokenType) Parser {
//	return func(input TokenStream) ParserResult {
//		it := input
//		for _, tt := range ts {
//			res := ExpectToken(tt)(it)
//			if res.Result == nil {
//				return Fail(input)
//			}
//			it = res.RemainingInput
//		}
//		return ParserResult{
//			Result:
//		}
//	}
//}

var Fail Parser = func(input TokenStream) ParserResult {
	return ParserResult{nil, input}
}
