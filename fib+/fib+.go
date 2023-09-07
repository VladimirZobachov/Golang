package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	result := fibonacci(n - 1)
	next := result[len(result)-1] + result[len(result)-2]
	return append(result, next)
}

func main() {
	fmt.Println(fibonacci(10))
}
