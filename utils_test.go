package top_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseString(t *testing.T) {
	b := []byte("aaa   bb cccc")

	expected := []string{"aaa", "bb", "cccc"}
	result := ParseString(b)

	assert.Equal(t, 3, len(result))
	assert.EqualValues(t, expected, result)
}
