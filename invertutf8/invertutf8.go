package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	inverted := invertString(input)
	fmt.Println("Inverted string:", inverted)
}

func invertString(input string) string {
	runes := []rune(input)
	length := utf8.RuneCountInString(input)
	inverted := make([]rune, length)
	for i, r := range runes {
		inverted[len(inverted)-i-1] = r
	}
	return string(inverted)
}
