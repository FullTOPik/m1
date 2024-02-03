package main

import "fmt"

func RemoveItemFromSlice(array []interface{}, index uint64) []interface{} {
	return append(array[:index], array[index+1:]...)
}

func main() {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var index uint64
	fmt.Printf("Enter index item for delete: ")
	fmt.Scan(&index)

	slice = RemoveItemFromSlice(slice, index)

	fmt.Println(slice...)

}
