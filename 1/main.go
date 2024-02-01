package main

//Дана структура Human (с произвольным набором полей и методов). 
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
)

type Human struct {
	name      string
	lastName  string
	firstName string
	age       uint16
}

type Action struct {
	Human
}

func (h *Human) GetFullName() string {
	return fmt.Sprintf("%s %s %s", h.firstName, h.name, h.lastName)
}

func (h *Human) SetAge(value uint16) *Human {
	h.age = value
	return h
}

func (h *Human) SetName(value string) *Human {
	h.name = value
	return h
}

func (h *Human) SetLastName(value string) *Human {
	h.lastName = value
	return h
}

func (h *Human) SetFirstName(value string) *Human {
	h.firstName = value
	return h
}

func NewHuman(name string, lastName string, firstName string, age uint16) *Human {
	return &Human{name: name, lastName: lastName, firstName: firstName, age: age}
}

func NewAction(name string, lastName string, firstName string, age uint16) *Action {
	return &Action{Human: Human{name: name, lastName: lastName, firstName: firstName, age: age}}
}

func main() {
	human := NewHuman("Ivan", "Ivanovich", "Ivanov", 35)
	action := NewAction("Gleb", "Sergeevich", "Kolonov", 16)

	humanFullName := human.GetFullName()
	fmt.Println(humanFullName)

	human.SetLastName("Cakovich").SetName("Cake")
	secondHumanFullName := human.GetFullName()
	fmt.Println(secondHumanFullName)

	actionFullName := action.GetFullName()
	fmt.Println(actionFullName)
}
