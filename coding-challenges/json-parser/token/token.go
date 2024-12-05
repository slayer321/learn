package token

type Type string

const (
	STRING             Type = "STRING"
	OPENCURLYBRACKET   Type = "{"
	CLOSEDCURLYBRACKET Type = "}"
	COMMA              Type = ","
	COLON              Type = ":"
)

type Token struct {
	Type  Type
	Value string
}

type TokenIterator struct {
	Tokens []Token
	Index  int
}

func NewTokenIterator(tokens []Token) TokenIterator {
	return TokenIterator{
		Tokens: tokens,
		Index:  0,
	}
}

// Peek returns the current index value
func (t *TokenIterator) Peek() string {
	return t.Tokens[t.Index].Value
}

func (t *TokenIterator) Next() {
	t.Index += 1
}

func (t *TokenIterator) HasNext() bool {
	if t.Index+1 == len(t.Tokens) {
		return false
	}
	return true
}
