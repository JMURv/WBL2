package main

import "fmt"

//В общем смысле, фасад работает как оболочка вокруг сложной системы,
//предоставляя простой интерфейс для взаимодействия с основной системой.

type StructA struct{}
type StructB struct{}

func (s *StructA) ActionA() {
	fmt.Println("Work A")
}

func (s *StructB) ActionB() {
	fmt.Println("Work B")
}

// Фасад - объединяющий интерфейс
type Facade struct {
	structA *StructA
	structB *StructB
}

func NewFacade() *Facade {
	return &Facade{
		structA: &StructA{},
		structB: &StructB{},
	}
}

func (f *Facade) MainAction() {
	f.structA.ActionA()
	f.structB.ActionB()
}

func main() {
	facade := NewFacade()
	facade.MainAction()
}
