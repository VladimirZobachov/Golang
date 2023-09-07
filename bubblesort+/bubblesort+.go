package main

import "fmt"

func bubbleSort(arr []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSort(arr, n-1)
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Array before sorting:", arr)

	bubbleSort(arr, len(arr))
	fmt.Println("Array after sorting:", arr)
}
