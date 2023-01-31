package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	strLen := utf8.RuneCountInString(text)
	padding := (width - strLen) / 2
	fmt.Println(padding, strLen)
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	fmt.Println(isPalindrome("hallelujah"))
	fmt.Println(isPalindrome("gogo"))
	fmt.Println(isPalindrome("gog"))
}

func isPalindrome(s string) bool {
	rs := []rune(s) // get slice of runes out of s
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}
