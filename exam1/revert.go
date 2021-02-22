package exam1

import (
	"strings"
)

func revertSentence(input string) (output string) {
	strs := strings.Split(input, " ")

	output = ""
	length := len(strs) - 1
	for idx, str := range strs {
		output += revertString(str)
		if idx < length {
			output += " "
		}
	}

	return output
}

func revertString(input string) (output string) {
	res := []rune(input)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	output = string(res)
	return output
}
