package neuralfunction

import (
	"github.com/jonysugianto/mathlib/matrix"
)

func AbsPower2(targets *matrix.Vector, outputs *matrix.Vector) *matrix.Vector {
	var e = outputs.Sub(targets)
	return e
}

func CrossEntropy(targets *matrix.Vector, outputs *matrix.Vector) *matrix.Vector {
	var logouts = outputs.Log()
	logouts.MulI(targets)
	return logouts
}
