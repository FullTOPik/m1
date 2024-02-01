package main

import (
	"fmt"
	"sync"
)

func GetArraySquires(numbers []int) []int {
	result := make([]int, len(numbers))
	wg := sync.WaitGroup{}

	for i, n := range numbers {
		wg.Add(1)

		go func(index int, number int) {
			defer wg.Done()

			result[index] = number * number
		}(i, n)
	}
	wg.Wait()

	return result
}

func TransformArrayToString(numbers []int) string {
	 result := "["

	for _, number := range numbers {
		result += fmt.Sprintf("%d, ", number)
	}

	result += "]"

	return result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	if len(numbers) == 0 {
		fmt.Println("Array is empty!")
		return
	}

	squires := GetArraySquires(numbers)
	result := TransformArrayToString(squires)

	fmt.Println(result)
}
