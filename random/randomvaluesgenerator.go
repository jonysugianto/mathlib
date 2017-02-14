package random

import (
	"math"

	"github.com/jonysugianto/mathlib/matrix"
)

func RandomValues(gen func() float64, size int) []float64 {
	var ret []float64
	for i := 0; i < size; i++ {
		ret = append(ret, gen())
	}
	return ret
}

func RandomValues2(gen func(float64) float64, size int, genparam float64) []float64 {
	var ret []float64
	for i := 0; i < size; i++ {
		ret = append(ret, gen(genparam))
	}
	return ret
}

// input vector 0 or 1
func AddNoiseBinaryVector(inputori *matrix.Vector, noise_level float64) *matrix.Vector {
	var input = inputori.Clone()
	var numberflipped = int(math.Ceil(noise_level * float64(input.Size())))
	var size = input.Size()
	for i := 0; i < numberflipped; i++ {
		var rnnr = Random(size)
		input.Values[rnnr] = 1 - input.Values[i]
	}
	return input
}
