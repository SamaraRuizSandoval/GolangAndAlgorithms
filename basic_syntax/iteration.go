package basic_syntax

import "strings"

const repeatCount = 5

func Repeat(character string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		//Strings in Go are immutable, meaning every concatenation involves copying memory to accommodate the new string.
		result += character
	}
	return result
}

func ImprovedRepeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
