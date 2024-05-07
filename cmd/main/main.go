package main

import (
	"fmt"
	"github.com/Abdirahman04/datumizer/intka"
	"github.com/Abdirahman04/datumizer/stringka"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("RandomString => ", stringka.RandomString(6))
	fmt.Println("RandomInt => ", intka.RandomInt(4))
}
