package main

import (
	"fmt"
	"strings"
)

func CheckString(str string) bool {
	newStr := strings.ToLower(str)
	runes := make(map[rune]bool, len(str))

	for _, r := range newStr {
		if runes[r] {
			return false
		}
		runes[r] = true
	}
	return true
}

func main() {
	fmt.Printf("Enter a string: ")
	var str string

	fmt.Scan(&str)

	result := CheckString(str)

	if result {
		fmt.Println("All characters in the string are unique")
	} else {
		fmt.Println("Characters are not unique!")
	}
}
