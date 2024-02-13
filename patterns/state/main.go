package main

import "fmt"

type State interface {
	performAction(r *Robot)
}

// Состояние "Включен" и "Выключен"
type OnState struct{}
type OffState struct{}
type WorkingState struct{ Name string }

func (s *OnState) performAction(r *Robot) {
	fmt.Printf("Робот %v включен\n", r.Name)
}

func (s *OffState) performAction(r *Robot) {
	fmt.Printf("Робот %v выключен\n", r.Name)
}

func (s *WorkingState) performAction(r *Robot) {
	fmt.Printf("Робот %v выполняет действие: %v\n", r.Name, s.Name)
}

// Сам робот
type Robot struct {
	Name  string
	state State
}

func (r *Robot) setState(s State) {
	r.state = s
}

func (r *Robot) performAction() {
	r.state.performAction(r)
}

func main() {
	robot := &Robot{Name: "John Doe"}
	robot.setState(&OnState{})
	robot.performAction()

	// Назначаем роботу работу и переводим в другое состояние
	robot.setState(&WorkingState{Name: "Cleaning"})
	robot.performAction()

	// Переключаем состояние на "выключен"
	robot.setState(&OffState{})
	robot.performAction()
}
