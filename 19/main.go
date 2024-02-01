package main

import "fmt"

func revertString(s string) string {
	runeArray := []rune(s)

	lenString := len(s) - 1
	for i, j := lenString, 0; i >= j; i, j = i-1, j+1 {
		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
	}

	return string(runeArray)
}

func revertString2(str string) string {
	var result string

	for _, c := range str {
		result += string(c)
	}

	return string(result)
}

func main() {
	s := "123456789"
	fmt.Println(s)

	result := revertString(s)
	fmt.Println(result)

	result = revertString2(s)
	fmt.Println(result)
}
