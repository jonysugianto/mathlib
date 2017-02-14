package conv

import (
	"github.com/jonysugianto/mathlib/matrix"
)

func ComputeConvOutputSize(input_length int, conv_length int, stride int) int {
	var conv_outputsize = int(float64(input_length-conv_length+1) / float64(stride))
	return conv_outputsize
}

func ConvInputsFromMatrixFloat(raw_input *matrix.MatrixFloat,
	input_w int, input_h int, kernel_w int, kernel_h int,
	stride_w int, stride_h int) *matrix.MatrixVector {
	var targetlength_x = ComputeConvOutputSize(input_w, kernel_w, stride_w)
	var targetlength_y = ComputeConvOutputSize(input_h, kernel_h, stride_h)
	var vectorsize = kernel_w * kernel_h
	var ret = matrix.CreateMatrixVector(targetlength_x, targetlength_y, vectorsize)

	for i := 0; i < targetlength_x; i++ {
		for j := 0; j < targetlength_y; j++ {
			var tileInput []float64
			for k := 0; k < kernel_w; k++ {
				for l := 0; l < kernel_h; l++ {
					tileInput = append(tileInput, raw_input.Get(i*stride_w+k, j*stride_h+l))
				}
			}
			var tileInputVector = matrix.CreateVector2(tileInput)
			ret.Put(i, j, tileInputVector)
		}
	}
	return ret
}

func ConvInputsFromListOfMatrixFloat(raw_inputs []*matrix.MatrixFloat,
	input_w int, input_h int, kernel_w int, kernel_h int,
	stride_w int, stride_h int) *matrix.MatrixVector {
	var targetlength_x = ComputeConvOutputSize(input_w, kernel_w, stride_w)
	var targetlength_y = ComputeConvOutputSize(input_h, kernel_h, stride_h)
	var vectorsize = kernel_w * kernel_h
	var ret = matrix.CreateMatrixVector(targetlength_x, targetlength_y, vectorsize)

	var number_inputs = len(raw_inputs)
	for i := 0; i < targetlength_x; i++ {
		for j := 0; j < targetlength_y; j++ {
			var tileInput []float64
			for k := 0; k < kernel_w; k++ {
				for l := 0; l < kernel_h; l++ {
					for ni := 0; ni < number_inputs; ni++ {
						tileInput = append(tileInput, raw_inputs[ni].Get(i*stride_w+k, j*stride_h+l))
					}
				}
			}
			var tileInputVector = matrix.CreateVector2(tileInput)
			ret.Put(i, j, tileInputVector)
		}
	}
	return ret
}

func ConvInputsFromMatrixVector(raw_input *matrix.MatrixVector,
	input_w int, input_h int, kernel_w int, kernel_h int,
	stride_w int, stride_h int) *matrix.MatrixVector {
	var targetlength_x = ComputeConvOutputSize(input_w, kernel_w, stride_w)
	var targetlength_y = ComputeConvOutputSize(input_h, kernel_h, stride_h)
	var vectorsize = raw_input.Values[0][0].Size() * kernel_w * kernel_h
	var ret = matrix.CreateMatrixVector(targetlength_x, targetlength_y, vectorsize)

	for i := 0; i < targetlength_x; i++ {
		for j := 0; j < targetlength_y; j++ {
			var tileInput []float64
			for k := 0; k < kernel_w; k++ {
				for l := 0; l < kernel_h; l++ {
					tileInput = append(tileInput, raw_input.Get(i*stride_w+k, j*stride_h+l).Values...)
				}
			}
			var tileInputVector = matrix.CreateVector2(tileInput)
			ret.Put(i, j, tileInputVector)
		}
	}
	return ret
}
