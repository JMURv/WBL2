package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // Получаем <nil> - нулевое значение, но тип - *os.PathError
	fmt.Println(err == nil) // Тип err - *fs.PathError, пусть и значение nil. nil в данном сравнении указывает на пустую ячейку в памяти, а err - на nil
}

// Вывод: <nil> false
