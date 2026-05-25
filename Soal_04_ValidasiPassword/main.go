package main

import (
	"fmt"
	"unicode"
)

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func main() {
	fmt.Println("=== Test IsValidPassword ===")
	fmt.Println(IsValidPassword("Abcdefg1"))
	fmt.Println(IsValidPassword("abcdefg1"))
	fmt.Println(IsValidPassword("ABCDEFG1"))
	fmt.Println(IsValidPassword("Abcdefgh"))
	fmt.Println(IsValidPassword("Ab1"))
}
