package scripts

import (
	"math/rand/v2"
	"strconv"
)

func RandomNumber() string {
	min := 1000
	max := 9999

	return strconv.Itoa(rand.IntN(max-min+1) + min)
}
