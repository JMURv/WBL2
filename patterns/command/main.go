package main

import "fmt"

// Контракт, которому команда должна удовлетворять
type Command interface {
	Execute()
}

// Receiver - получатель команды
type Light struct {
	isOn bool
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("on")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("off")
}

// Команда для включения/выключения света
type TurnOnCommand struct {
	light *Light
}

type TurnOffCommand struct {
	light *Light
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Структура, запускающая команды
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := NewLight()
	remote := &RemoteControl{}

	remote.SetCommand(NewTurnOnCommand(light))
	remote.PressButton()

	remote.SetCommand(NewTurnOffCommand(light))
	remote.PressButton()
}
