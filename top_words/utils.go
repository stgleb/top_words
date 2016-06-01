package top_words

func split(s string) []string {
	arraySize := 1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			arraySize++
		}

		for s[i] == ' ' {
			i++
		}
	}
	array := make([]string, arraySize)

	currentStrInd := 0
	currentStr := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(currentStr) > 0 {
				array[currentStrInd] = currentStr
				currentStrInd++
				currentStr = ""
			}
		} else {
			currentStr += string(s[i])
		}
	}

	if len(currentStr) > 0 {
		array[arraySize-1] = currentStr
	}
	return array[:]
}

func ParseString(bytes []byte) []string {
	s := string(bytes)

	return split(s)
}
