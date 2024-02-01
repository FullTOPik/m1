package main

import (
	"fmt"
	"strings"
)

func revertWordsInString(str string) string {
	stringArray := strings.Split(str, " ")

	for i, j := 0, len(stringArray)-1; i <= j; i, j = i+1, j-1 {
		stringArray[i], stringArray[j] = stringArray[j], stringArray[i]
	}

	return strings.Join(stringArray, " ")
}

func revertWordsInString2(str string) string {
	stringArray := strings.Fields(str)

	var b strings.Builder

	for i := len(stringArray) - 1; i >= 0; i-- {
		b.WriteString(stringArray[i])
		b.WriteString(" ")
	}

	return b.String()
}

func main() {
	s := "dog cat warm"
	fmt.Println(revertWordsInString(s))
	fmt.Println(revertWordsInString2(s))
}
