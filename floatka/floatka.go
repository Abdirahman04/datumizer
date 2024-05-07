package floatka

import (
	"math"
	"math/rand"
)

func RandomFloat(ln, dec int) float64 {
	whole := rand.Intn(ln)
	decimals := rand.Intn(dec)

	num := whole + decimals

	fl := math.Pow(10, float64(dec))

	return float64(num) / fl
}
