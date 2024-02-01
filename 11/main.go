package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func NewMap(arr []int) map[int]bool {
	set := make(map[int]bool)
	for _, value := range arr {
		set[value] = true
	}

	return set
}

func main() {
	firstSet := []int{1, 2, 3, 4, 5, 6, 7}
	secondSet := []int{5, 6, 7, 3, 8, 12}

	firstMap := NewMap(firstSet)
	secondMap := NewMap(secondSet)

	result := make(map[int]bool)

	for key := range firstMap {
		if secondMap[key] {
			result[key] = true
		}
	}

	for key := range result {
		fmt.Printf("%d, ", key)
	}
}
