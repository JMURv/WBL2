package main

import "fmt"

// Product - контракт, который должны реализовывать продукты
type Product interface {
	GetName() string
}

// ProductA и ProductB - реализации различных продуктов
// Пока продукты реализуют один и тот же интерфейс Product, они отлично вписываются в бизнес-логику и всё работает
type ProductA struct{ Name string }
type ProductB struct{ Name string }

func (p *ProductA) GetName() string {
	return p.Name
}

func (p *ProductB) GetName() string {
	return p.Name
}

// Creator - контракт для фабричных структур
type Creator interface {
	CreateProduct() Product
}

// CreatorA и CreatorB - реализации фабрик для ПродуктаА и ПродуктаБ
type CreatorA struct{}
type CreatorB struct{}

func (c *CreatorA) CreateProduct() Product {
	return &ProductA{Name: "Product A"}
}

func (c *CreatorB) CreateProduct() Product {
	return &ProductB{Name: "Product B"}
}

// Смысл паттерна в том, чтобы отвязаться от конкретных реализаций продукта и чтобы наращивать новые было проще
func main() {
	creatorA := &CreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Println(productA.GetName())

	creatorB := &CreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Println(productB.GetName())
}
