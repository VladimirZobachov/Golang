package main

import (
	"fmt"
	"strconv"
)

func check(count int) bool {
	reversed := 0
	original := count

	for count > 0 {
		digit := count % 10
		reversed = reversed*10 + digit
		count = count / 10
	}

	return original == reversed
}

func checkConvert(count int) bool {

	str := strconv.Itoa(count)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(checkConvert(5234725))
	fmt.Println(check(5234725))
}
