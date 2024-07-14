package main

import (
	"fmt"
	"unicode"
)

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	i, j := 0, n-1

	for i < j {
		if !unicode.IsLetter(runes[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(runes[j]) {
			j--
			continue
		}
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func main() {
	input := "NEGIE1"
	output := reverse(input)
	fmt.Println(output)
}
