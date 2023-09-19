package main

import "fmt"

func numunicchars(str string) map[string]int {
	result := map[string]int{}
	for _, runeValue := range str {
		result[string(runeValue)] = result[string(runeValue)] + 1
	}
	return result
}

func main() {
	fmt.Println(numunicchars("aaaabbbeeerrr"))
}
