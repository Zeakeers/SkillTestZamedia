package main

import "fmt"

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Test IsPalindrome ===")
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("hello"))
	fmt.Println(IsPalindrome("madam"))
	fmt.Println(IsPalindrome("Racecar"))
}
