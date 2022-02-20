package main

import (
	"fmt"

	"github.com/DShomin/lomitcoin/person"
)

func main() {
	lomit := person.Person{}
	lomit.SetDetails("lomit", 30)
	fmt.Println(lomit)
}
