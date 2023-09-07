package main

// import "fmt"

// func fact(num int) int {
// 	if num == 0 {
// 		return 1
// 	}

// 	result := 1

// 	for i := num; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

func main() {
	//fmt.Println(fact(5))
	str := "абв"

	for _, r := range str {
		println(r, string(r))
	}
}
