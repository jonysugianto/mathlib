package matrix

import (
	"fmt"
	"math"
)

type Vector struct {
	Values []float64
}

func CreateVector(size int) *Vector {
	var ret = new(Vector)
	ret.Values = make([]float64, size)
	return ret
}

func CreateVector2(vs []float64) *Vector {
	var ret = CreateVector(len(vs))
	ret.Copy(vs)
	return ret
}

func (this *Vector) Size() int {
	return len(this.Values)
}

func (this *Vector) Clone() *Vector {
	var retvalues = make([]float64, len(this.Values))
	copy(retvalues, this.Values)
	var ret = new(Vector)
	ret.Values = retvalues
	return ret
}

func (this *Vector) Copy(data []float64) {
	copy(this.Values, data)
}

func (this *Vector) Log() *Vector {
	var ret = CreateVector(this.Size())
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = math.Log(this.Values[i])
	}
	return ret
}

func (this *Vector) LogI() {
	var size = this.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = math.Log(this.Values[i])
	}
}

func (this *Vector) Square() *Vector {
	var ret = CreateVector(this.Size())
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] * this.Values[i]
	}
	return ret
}

func (this *Vector) SquareI() {
	var size = this.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] * this.Values[i]
	}
}

func (this *Vector) Add(v *Vector) *Vector {
	var ret = CreateVector(v.Size())
	var size = v.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] + v.Values[i]
	}
	return ret
}

func (this *Vector) AddI(v *Vector) {
	var size = v.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] + v.Values[i]
	}
}

func (this *Vector) AddScalar(s float64) *Vector {
	var ret = CreateVector(this.Size())
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] + s
	}
	return ret
}

func (this *Vector) AddScalarI(s float64) {
	var size = this.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] + s
	}
}

func (this *Vector) Sub(v *Vector) *Vector {
	var ret = CreateVector(v.Size())
	var size = v.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] - v.Values[i]
	}
	return ret
}

func (this *Vector) SubI(v *Vector) {
	var size = v.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] - v.Values[i]
	}
}

func (this *Vector) SubScalar(s float64) *Vector {
	var ret = CreateVector(this.Size())
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] - s
	}
	return ret
}

func (this *Vector) SubScalarI(s float64) {
	var size = this.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] - s
	}
}

func (this *Vector) Mul(v *Vector) *Vector {
	var ret = CreateVector(v.Size())
	var size = v.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] * v.Values[i]
	}
	return ret
}

func (this *Vector) MulI(v *Vector) {
	var size = v.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] * v.Values[i]
	}
}

func (this *Vector) MulScalar(s float64) *Vector {
	var ret = CreateVector(this.Size())
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret.Values[i] = this.Values[i] * s
	}
	return ret
}

func (this *Vector) MulScalarI(s float64) {
	var size = this.Size()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] * s
	}
}

func (this *Vector) Sum() float64 {
	var size = this.Size()
	var sum float64 = 0
	for i := 0; i < size; i++ {
		sum = sum + this.Values[i]
	}
	return sum
}

func (this *Vector) Magnitude() float64 {
	var size = this.Size()
	var sum float64 = 0
	for i := 0; i < size; i++ {
		sum = sum + this.Values[i]*this.Values[i]
	}
	return math.Sqrt(sum)
}

func (this *Vector) MakeUnity() {
	var size = this.Size()
	var mag = this.Magnitude()
	for i := 0; i < size; i++ {
		this.Values[i] = this.Values[i] / mag
	}
}

func (this *Vector) String() string {
	var ret string
	var size = this.Size()
	for i := 0; i < size; i++ {
		ret = ret + " " + fmt.Sprintf("%.2f", this.Values[i])
	}
	return ret
}

func (this *Vector) AsMatrixFloat(row int, col int) *MatrixFloat {
	var ret = CreateMatrixFloat(row, col)
	var index = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ret.Values[i][j] = this.Values[index]
			index++
		}
	}
	return ret
}
