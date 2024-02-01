package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации.

// var justString string Глобальная переменная

func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func someFunc() string {
	v := createHugeString(1 << 10)

	temp := make([]rune, 100)
	copy(temp, []rune(v)[:100])
	justString := string(temp)

	return justString
}

func main() {
	var justString string
	justString = someFunc()
	fmt.Println(justString)
}
