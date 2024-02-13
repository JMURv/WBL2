package main

import "fmt"

// Strategy - интерфейс для стратегии
type Strategy interface {
	DoSmth()
}

// ConcreteStrategyA и ConcreteStrategyB - конкретные реализации стратегий
type StrategyA struct{}
type StrategyB struct{}

func (s *StrategyA) DoSmth() {
	fmt.Println("DoSmth with StrategyA")
}

func (s *StrategyB) DoSmth() {
	fmt.Println("DoSmth with StrategyB")
}

// Context - контекст, который использует стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) Execute() {
	c.strategy.DoSmth()
}

func main() {
	context := &Context{}

	context.SetStrategy(&StrategyA{})
	context.Execute()

	context.SetStrategy(&StrategyB{})
	context.Execute()
}
