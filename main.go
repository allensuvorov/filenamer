package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter LeetCode problem name to produce a nicely formatted file name: ")
	if text, err := reader.ReadString('\n'); err == nil {
		fmt.Println(formatFileName(text))
	}
}

func formatFileName(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = underScore(s)
	s = s + ".go"
	return s
}

func underScore(s string) string {
	b := []rune{}
	inSymbols := false

	for _, v := range s {

		if !inSymbols {
			if isSymbol(v) { // enter symbols
				b = append(b, '_')
				inSymbols = true
			} else {
				b = append(b, v)
			}
		} else {
			if !isSymbol(v) { // exit symbols
				b = append(b, v)
				inSymbols = false
			}
		}
	}
	return string(b)
}

func isSymbol(r rune) bool {
	if unicode.IsDigit(r) || unicode.IsLetter(r) {
		return false
	}
	return true
}
