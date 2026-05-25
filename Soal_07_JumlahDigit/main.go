package main

import "fmt"

func DigitSum(n int) int {
	if n < 0 {
		n = -n
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println("=== Test DigitSum ===")
	fmt.Println(DigitSum(123))
	fmt.Println(DigitSum(9999))
	fmt.Println(DigitSum(0))
	fmt.Println(DigitSum(-456))
}
