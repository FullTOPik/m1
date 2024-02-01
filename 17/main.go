package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
func main() {
	array := []int{1, 2, 7, 9, 12, 23, 34, 54, 65, 88, 97}
	target := 9
	fmt.Printf("Array = %v\nTarget = %d, index = %d", array, target, BinarySearch(array, target))
}

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		middle := arr[mid]
		if middle < target {
			left = mid + 1
		} else if middle > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}