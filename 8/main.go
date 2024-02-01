package main

import (
	"fmt"
	"os"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var number uint64
	var i uint8

	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &number); err != nil {
		fmt.Println("Недопустимый ввод")
		return
	}
	
	fmt.Println("Введите номер бита, что необходимо инвертировать")
	if _, err := fmt.Fscan(os.Stdin, &i); err != nil {
		fmt.Println("Недопустимый ввод")
		return
	}

	if i > 63 {
		fmt.Println("Недопустимое количество")
		return
	}

	fmt.Printf("Введённое десятичное число: %d\nБитовая запись числа: %064b\n", number, number)
	number ^= 1 << i
	fmt.Printf("Полученное десятичное число: %d\nБитовая запись числа: %064b", number, number)
}
