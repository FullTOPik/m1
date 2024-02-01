package main

import "fmt"

func main() {
	firstVar := 1
	secondVar := 2

	fmt.Printf("FirstVar: %d, SecondVar: %d", firstVar, secondVar)

	firstVar, secondVar = secondVar, firstVar
	fmt.Printf("\nFirstVar: %d, SecondVar: %d", firstVar, secondVar)
	return
}
