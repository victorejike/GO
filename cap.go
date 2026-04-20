package main

import (
	"strings"
)

func capitalized(text string) string {
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return result
}
func transforma(text string) string {
	tokens := strings.Fields(text)
	var result []string

	for i := 0; i < len(tokens); i++ {
		tokens := tokens[i]

		if tokens == "(cap)" && len(result) > 0 {
			result[len(result)-1] = capitalized(result[len(result)-1])
			continue
		}

		if tokens == "(up)" && len(result) > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}

	}
}
