package neuralfunction

import (
	"math"
)

func Sigmoid(x float64) float64 {
	return (1.0 / (1.0 + math.Pow(math.E, -x)))
}

func SigmoidArray(xs []float64) []float64 {
	var size = len(xs)
	var ret = make([]float64, size)
	for i := 0; i < size; i++ {
		ret[i] = Sigmoid(xs[i])
	}
	return ret
}

func SoftmaxArray(xs []float64) []float64 {
	var size = len(xs)
	var sum = float64(0)
	for i := 0; i < size; i++ {
		sum += math.Exp(xs[i])
	}

	var ret = make([]float64, size)
	for i := 0; i < size; i++ {
		ret[i] = math.Exp(xs[i]) / sum
	}
	return ret
}

func SigmoidSoftmaxArray(xs []float64) []float64 {
	var temp = SigmoidArray(xs)
	return SoftmaxArray(temp)
}

func Tanh(x float64) float64 {
	return (2 * Sigmoid(x)) - 1
}

func TanhArray(xs []float64) []float64 {
	var size = len(xs)
	var ret = make([]float64, size)
	for i := 0; i < size; i++ {
		ret[i] = Tanh(xs[i])
	}
	return ret
}
