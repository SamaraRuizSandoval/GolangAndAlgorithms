package basic_syntax

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	result := ""
	for i := 0; i < repeatCount; i++ {
		//Strings in Go are immutable, meaning every concatenation involves copying memory to accommodate the new string.
		result += character
	}
	return result
}

func ImprovedRepeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
