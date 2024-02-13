package main

import "fmt"

// Студенты разных типов, над ними нужно провести аудит. Мы не хотим добавлять методы прямо в структуры, а хотим вынести эту логику в Visitor'а
// Студенты бакалавриата или магистратуры должны имплементировать метод Accept, чтобы принять Visitor'а
type Student interface {
	Accept(v Visitor)
}

type MagStudent struct {
	Name string
}

type BacStudent struct {
	Name string
}

func (s *BacStudent) Accept(v Visitor) {
	v.VisitBacStudent(s)
}

func (s *MagStudent) Accept(v Visitor) {
	v.VisitMagStudent(s)
}

// Контракт для посетителя. Посетитесь имлементирует метод Visit для каждого типа, к которому надо постучаться
type Visitor interface {
	VisitBacStudent(s *BacStudent)
	VisitMagStudent(s *MagStudent)
}

type Auditor struct{}

func (a *Auditor) VisitBacStudent(s *BacStudent) {
	fmt.Println("Аудит для бакалавра:", s.Name)
}

func (a *Auditor) VisitMagStudent(s *MagStudent) {
	fmt.Println("Аудит для магистра:", s.Name)
}

// Структура, содержащая студентов
type University struct {
	Students []Student
}

func (u *University) AddStudent(student Student) {
	u.Students = append(u.Students, student)
}

func (u *University) ConductAudit(v Visitor) {
	for _, s := range u.Students {
		s.Accept(v)
	}
}

func main() {
	u := &University{}

	u.AddStudent(&BacStudent{Name: "John"})
	u.AddStudent(&MagStudent{Name: "Doe"})
	u.ConductAudit(&Auditor{})
}
