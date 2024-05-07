package floatka

import (
	"math"
	"math/rand"
)

func RandomFloat(ln, dec float64) float64 {
	whole := rand.Intn(int(math.Pow(10, ln)))
	decimals := rand.Intn(int(math.Pow(10, dec)))

	num := whole + decimals

	fl := math.Pow(10, dec)

	return float64(num) / fl
}
