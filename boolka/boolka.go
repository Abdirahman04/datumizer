package boolka

import "math/rand"

func RandomBool() bool {
	return rand.Intn(2) <= 0
}
