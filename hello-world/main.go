package main

import (
	"fmt"
)

// simple struct type with method
type Rect struct {
	Width  int
	Height int
}

func (r *Rect) Area() int {
	return r.Width * r.Height
}

// simple array type with a method
type Rects []Rect

func (rs Rects) Area() int {
	sumArea := 0

	for _, r := range rs {
		sumArea += r.Area()
	}
	return sumArea
}

// function callback
type Foo func(x int) int

func (cb Foo) Add(x int) int {
	return cb(x) + x
}

// interface usage
type Shaper interface {
	Area() int
}

func Describe(s Shaper) {
	fmt.Println("Area of Shape:", s.Area())
}

//sub types

type Person struct {
	Name string
	Address
}

func (p *Person) String() string {
	return p.Name + "-" + p.Address.String()
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (addr *Address) String() string {

	return addr.Street + "-" + addr.City + "-" + addr.Country
}

func main() {

	var callback Foo = func(x int) int {
		return x - 1
	}

	r1 := &Rect{Width: 5, Height: 120}
	r2 := &Rect{Width: 8, Height: 59}
	ls := Rects{*r1, *r2}

	p := &Person{
		Name: "Junaid",
		Address: Address{
			City:    "Lahore",
			Country: "Pakistan",
		},
	}

	fmt.Println("Hello World!")
	fmt.Println("Add using callback:", callback.Add(5))
	Describe(r1)
	Describe(ls)
	fmt.Println("Person Information:", p.String())
}
