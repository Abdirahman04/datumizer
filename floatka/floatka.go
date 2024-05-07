package floatka

import (
	"math"
	"math/rand"
)

func RandomFloat(ln, dec float64) float64 {
	dc := math.Pow(10, dec)
	whole := rand.Intn(int(math.Pow(10, ln)))
	decimals := rand.Intn(int(dc))

	num := (whole * int(dc)) + decimals

	return float64(num) / dc
}
