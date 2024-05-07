package intka

import "math/rand"

func RandomInt(ln int) int {
	x := 10
	for range ln {
		x *= 10
	}

	return rand.Intn(x)
}
