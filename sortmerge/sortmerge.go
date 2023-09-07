package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		merged[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}

	return merged
}

func main() {
	arr1 := []int{3, 1, 2, 5, 4}
	arr2 := []int{9, 7, 8, 6, 10}

	sorted1 := mergeSort(arr1)
	sorted2 := mergeSort(arr2)

	fmt.Println("Sorted Array 1:", sorted1)
	fmt.Println("Sorted Array 2:", sorted2)
}
