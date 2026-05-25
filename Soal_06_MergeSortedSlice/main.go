package main

import "fmt"

func MergeSorted(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

func main() {
	fmt.Println("=== Test MergeSorted ===")
	fmt.Println(MergeSorted([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println(MergeSorted([]int{1, 5, 9}, []int{2, 3, 7, 10}))
}
