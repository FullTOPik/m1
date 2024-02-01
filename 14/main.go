package main

import "fmt"

func getType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func main() {
	var item interface{}

	item = 1
	fmt.Printf("%s\n", getType(item))

	item = ""
	fmt.Printf("%s\n", getType(item))
}
