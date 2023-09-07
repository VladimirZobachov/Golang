package main

import "fmt"

func fibonacci(num int) []int {

	result := make([]int, num)

	for i := 0; i < num; i++ {

		if i == 0 || i == 1 {
			result[i] = i
			continue
		}

		result[i] = result[i-1] + result[i-2]
	}

	return result

}

func main() {
	fmt.Println(fibonacci(10))
}
