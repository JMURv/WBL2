package main

import "fmt"

// Handler - контракт, который должны реализовывать обработчики
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// HandlerA - реализация обработчика A
type HandlerA struct {
	nextHandler Handler
}

func (c *HandlerA) SetNext(handler Handler) {
	c.nextHandler = handler
}

func (c *HandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("HandlerA handled the request")
	} else if c.nextHandler != nil {
		c.nextHandler.Handle(request)
	}
}

// HandlerB - реализация обработчика B
type HandlerB struct {
	nextHandler Handler
}

func (c *HandlerB) SetNext(handler Handler) {
	c.nextHandler = handler
}

func (c *HandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled the request")
	} else if c.nextHandler != nil {
		c.nextHandler.Handle(request)
	} else if c.nextHandler == nil {
		fmt.Println("No one can handle the request")
	}
}

func main() {
	handlerA := &HandlerA{}
	handlerB := &HandlerB{}

	handlerA.SetNext(handlerB)

	handlerA.Handle("A")
	handlerA.Handle("B")
	handlerA.Handle("C")
}
