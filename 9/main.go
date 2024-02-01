package main

import "fmt"

func main() {
	array := [100]int{}
	firstCh := make(chan int)
	secondCh := make(chan int)

	//заполнение масива значениями
	for i := range array {
		array[i] = i + 1
	}

	//пишем числа в первый канал
	go func() {
		for _, value := range array {
			firstCh <- value
		}
		close(firstCh)
	}()

	//умножаем на 2 и пишем во второй канал
	go func() {
		for value := range firstCh {
			secondCh <- value * value
		}
		close(secondCh)
	}()

	for value := range secondCh {
		fmt.Println(value)
	}
}
