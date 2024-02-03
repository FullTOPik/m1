package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	var firstNumber big.Int

	fmt.Printf("Enter a first number: ")
	if _, err := fmt.Scanln(&firstNumber); err != nil {
		fmt.Println("Invalid format!")
	}

	var secondNumber big.Int
	fmt.Printf("Enter a second number: ")
	if _, err := fmt.Scanln(&secondNumber); err != nil {
		fmt.Println("Invalid format!")
	}

	fmt.Printf(`
Select action:
 1)Press 1 for add numbers
 2)Press 2 for sub numbers
 3)Press 3 for multiplication numbers
 4)Press 4 for sub numbers
	`)

	ch, _, err := keyboard.GetSingleKey()
	if err != nil {
		return
	}

	char, err := strconv.Atoi(string(ch))
	if err != nil {
		fmt.Println("Incorrect value!")
		return
	}

	switch char {
	case 1:
		fmt.Println(firstNumber.Add(&firstNumber, &secondNumber))
	case 2:
		fmt.Println(firstNumber.Sub(&firstNumber, &secondNumber))
	case 3:
		fmt.Println(firstNumber.Mul(&firstNumber, &secondNumber))
	case 4:
		fmt.Println(firstNumber.Div(&firstNumber, &secondNumber))
	}

}
