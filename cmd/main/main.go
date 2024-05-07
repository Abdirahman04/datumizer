package main

import (
	"fmt"
	"github.com/Abdirahman04/datumizer/boolka"
	"github.com/Abdirahman04/datumizer/dateka"
	"github.com/Abdirahman04/datumizer/floatka"
	"github.com/Abdirahman04/datumizer/intka"
	"github.com/Abdirahman04/datumizer/stringka"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("RandomString => ", stringka.RandomString(6))
	fmt.Println("RandomInt => ", intka.RandomInt(4))
	fmt.Println("RandomFloat => ", floatka.RandomFloat(3, 2))
	fmt.Println("RandomBool => ", boolka.RandomBool())
	fmt.Println("RandomDate => ", dateka.RandomDate())
}
