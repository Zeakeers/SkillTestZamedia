package main

import "fmt"

func CharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	return freq
}

func main() {
	fmt.Println("=== Test CharFrequency ===")
	result := CharFrequency("hello world")
	for ch, count := range result {
		fmt.Printf("'%c': %d\n", ch, count)
	}
}
