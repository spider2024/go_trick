package simple

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi my name is %s\n", p.Name)
}

// NewPerson simple factory
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

type person struct {
	Name string
	Age  int
}

type PersonExport interface {
	Greet()
}

func (p person) Greet() {
	fmt.Printf("Hi my name is %s\n", p.Name)
}

// NewPersonAbstract return an interface instead a struct
func NewPersonAbstract(name string, age int) PersonExport {
	return person{
		Name: name,
		Age:  age,
	}
}
