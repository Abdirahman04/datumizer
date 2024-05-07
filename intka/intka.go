package intka

import "math/rand"

func RandomInt(ln int) int {
	x := 1
	for range ln {
		x *= 10
	}

	return rand.Intn(x)
}
