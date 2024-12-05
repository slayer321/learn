package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tokens = []Token{
		{
			Type:  OPENCURLYBRACKET,
			Value: "{",
		},
		{
			Type:  STRING,
			Value: "key",
		},
		{
			Type:  COLON,
			Value: ":",
		},
		{
			Type:  STRING,
			Value: "value",
		},
		{
			Type:  CLOSEDCURLYBRACKET,
			Value: "}",
		},
	}
)

func TestPeek(t *testing.T) {

	Iterator := NewTokenIterator(tokens)

	assert.Equal(t, "{", Iterator.Peek())

}

func TestNext(t *testing.T) {

	Iterator := NewTokenIterator(tokens)

	assert.Equal(t, "{", Iterator.Peek())

	Iterator.Next()

	assert.Equal(t, "key", Iterator.Peek())

}

func TestHasNext(t *testing.T) {

	Iterator := NewTokenIterator(tokens)

	assert.Equal(t, "{", Iterator.Peek())

	Iterator.Next()
	Iterator.Next()
	assert.True(t, Iterator.HasNext())

	Iterator.Next()
	Iterator.Next()

	assert.Equal(t, "}", Iterator.Peek())
	assert.False(t, Iterator.HasNext())

}
