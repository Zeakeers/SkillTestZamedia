package main

import "fmt"

func IndexOf(numbers []int, target int) int {
	for i, v := range numbers {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("=== Test IndexOf ===")
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println(IndexOf(nums, 30))
	fmt.Println(IndexOf(nums, 99))
	fmt.Println(IndexOf(nums, 10))
}
