package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p *Person) Intro() string {
	return fmt.Sprintln(p.Name, p.Age, p.Email)
}

type TaxPayer struct {
	TaxType   string
	YearlyTax int32
}

func (tp *TaxPayer) Intro() string {
	return "TaxType: " + tp.TaxType + ", paying yearly: " + strconv.Itoa(int(tp.YearlyTax))
}

type Employee struct {
	*Person
	*TaxPayer
	Designation string
}

func (e *Employee) Intro() string {
	return e.Person.Intro() + ", I am working as " + e.Designation + e.TaxPayer.Intro()
}

func LearnStructs() {
	me := Employee{
		Person:      &Person{"junaid", 22, "junaid"},
		TaxPayer:    &TaxPayer{"Income tax 24 %% slab", 24},
		Designation: "Software Developer",
	}

	me.Email = "junaid.shakur.js@gmail.com"
	fmt.Println("My Info", me.Intro())
}
