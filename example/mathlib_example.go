package main

import (
	"fmt"

	"github.com/jonysugianto/mathlib/matrix"
	"github.com/jonysugianto/mathlib/random"
)

func main() {
	var vvalues1 = random.RandomValues2(random.GaussianScaleMinusScale, 3, 10)
	var v1 = matrix.CreateVector2(vvalues1)
	fmt.Println(v1)

	var vvalues2 = random.RandomValues(random.GaussianOneMinusOne, 3)
	var v2 = matrix.CreateVector2(vvalues2)
	fmt.Println(v2)

	var v3 = v1.Mul(v2)
	fmt.Println(v3)

	fmt.Println(v3.Magnitude())
	fmt.Println(v3.Sum())
}
