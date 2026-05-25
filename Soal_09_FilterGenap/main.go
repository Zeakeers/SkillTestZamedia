package main

import "fmt"

func FilterEven(numbers []int) []int {
	result := []int{}
	for _, v := range numbers {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("=== Test FilterEven ===")
	fmt.Println(FilterEven([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(FilterEven([]int{7, 11, 13}))
	fmt.Println(FilterEven([]int{2, 4, 6, 8}))
}
