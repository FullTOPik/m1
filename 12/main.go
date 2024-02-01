package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.

func main() {
	arrayStr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, value := range arrayStr {
		set[value] = true
	}

	for key := range set {
		fmt.Printf("%s, ", key)
	}
}
