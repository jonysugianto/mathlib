package random

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("init")
	rand.Seed(time.Now().UnixNano())
}

func Random(n int) int {
	return rand.Intn(n)
}

func RandomGaussian() float64 {
	return rand.Float64()
}

func BinomialOneMinusOne() float64 {
	if rand.Float64() < 0.5 {
		return 1
	} else {
		return -1
	}
}

func GaussianOneMinusOne() float64 {
	return BinomialOneMinusOne() * rand.Float64()
}

func GaussianScaleMinusScale(scale float64) float64 {
	return BinomialOneMinusOne() * rand.Float64() * scale
}
