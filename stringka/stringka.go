package stringka

import (
	"math/rand"
)

func RandomString(ln int) string {
	var str string

	chr := [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for range ln {
		str += string(chr[rand.Intn(26)])
	}

	return str
}
