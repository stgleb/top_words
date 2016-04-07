package tests


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/top_words/top_words"
)

func TestParseString(t *testing.T){
	b := []byte("aaa   bb cccc")

	expected := []string{"aaa", "bb", "cccc"}
	result := top_words.ParseString(b)

	assert.Equal(t, 3, len(result))
	assert.EqualValues(t, expected, result)
}
