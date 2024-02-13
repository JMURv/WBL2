package main

import "fmt"

// Продукт, который мы хотим построить
type Product struct {
	elements []string // Тут составные части продукта
}

// Добавляем какой-либо элемент к продукту (используется ConcreteBuilder'ами)
func (p *Product) AddElem(elem string) {
	p.elements = append(p.elements, elem)
}

func (p *Product) GetAllElems() {
	fmt.Println(p.elements)
}

// Builder - контракт для строителя, то есть те методы, которые должны реализовывать конкретные строители (ConcreteBuilder'ы)
type Builder interface {
	Build1()
	Build2()
	GetResult() *Product
}

// ConcreteBuilder - конкретная реализация строителя
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) Build1() {
	b.product.AddElem("Part 1")
}

func (b *ConcreteBuilder) Build2() {
	b.product.AddElem("Part 2")
}

func (b *ConcreteBuilder) GetResult() *Product {
	return b.product
}

// Director - прораб, управляющий строителем
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.Build1()
	d.builder.Build2()
}

func main() {
	builder := NewConcreteBuilder()  // Создаем экземпляр строителя, который будет передаваться в директора
	director := NewDirector(builder) // Директор принимает строителя и..

	director.Construct()           // Конструирует продукт
	product := builder.GetResult() // Тут мы его получаем

	product.GetAllElems() // Вызываем метод продукта
}
