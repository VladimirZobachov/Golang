package main

import (
	"fmt"
	"unicode/utf8"
)

func numchars(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	fmt.Println(numchars("выфафвпвф"))
}
