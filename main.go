package main

import (
	"fmt"
	"time"
)

// create class
type person struct {
	name string
	age  int
}

// create class method
func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d\n", p.name, p.koreanAge())
}

func (p person) koreanAge() int {
	currentYear := int(time.Now().Year())
	return currentYear - p.age + 1
}

func main() {
	lomit := person{name: "Lomit", age: 1993}
	lomit.sayHello()
}
