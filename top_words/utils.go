package top_words

import "strings"


func parseString(bytes []byte) []string {
	s := string(bytes)

	return strings.SplitAfter(s, " ")
}
