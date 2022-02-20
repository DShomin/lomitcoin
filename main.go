package main

import "fmt"

func main() {
	a := 2
	// copy value
	// b := a
	// memorry address save
	b := &a
	// change value a
	a = 50
	fmt.Println(*b, b, &a)
}
