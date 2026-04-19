package main

import (
	"fmt"
	"strings"
)

func capitalized(text string) string {
	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	return result
}

func transforma(text string) string {
	token := strings.Fields(text)
	var result []string

	for i := 0; i < len(result); i++ {
		token := token[i]

		if token == "(cap)" && len(result) > 0 {
			result[len(result)-1] = capitalized(result[len(result)-1])
			continue
		}

		if token == "(up)" && len(result) > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}

		if token == "(low)" && len(result) > 0 {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
			continue
		}
	}
	return result[len(text)]
}

func main() {
	fmt.Println(transforma("hello(cap) this my first name"))
}
