package math

import (
	"math/rand"
)

func RandInt64(min, max, seed int64) int64 {
	rand.Seed(seed)
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
