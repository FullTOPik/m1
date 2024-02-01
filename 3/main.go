package main

import (
	"fmt"
	"sync"
)


func GetSumSquires(array []int) int {
	var result int
	wg := sync.WaitGroup{}

	for _, n := range array {
		wg.Add(1)

		go func(number int) {
			result += number * number
			wg.Done()
		}(n)
	}
	wg.Wait()

	return result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	if len(numbers) == 0 {
		fmt.Println("Array is empty!")
		return
	}

	result := GetSumSquires(numbers)
	fmt.Println(result)
}