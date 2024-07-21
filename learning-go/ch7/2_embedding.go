package main

import "fmt"

func main2() {
	m := Manager{
		Employee: Employee{
			Name: "John Doe",
			ID:   "JD123",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)

	//var eFail Employee = m
	var eOk Employee = m.Employee
	fmt.Println(eOk)

	o2 := Outer2{
		Inner: Inner{
			X: 10,
		},
		S: "Hello",
	}
	fmt.Println(o2.Double()) // Inner: 20
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("Name: %s, ID: %s", e.Name, e.ID)
}

// Employee를 embedding한 Manager
type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return m.Reports
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

type Outer2 struct {
	Inner
	S string
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}
