package matrix

type MatrixFloat struct {
	Values [][]float64
}

type MatrixVector struct {
	Values [][]*Vector
}

func CreateMatrixFloat(rowsize int, colsize int) *MatrixFloat {
	var values [][]float64
	for i := 0; i < rowsize; i++ {
		var row = make([]float64, colsize)
		values = append(values, row)
	}
	var ret = new(MatrixFloat)
	ret.Values = values
	return ret
}

func (this *MatrixFloat) Size() int {
	return len(this.Values) * len(this.Values[0])
}

func (this *MatrixFloat) Get(r int, c int) float64 {
	return this.Values[r][c]
}

func (this *MatrixFloat) Put(r int, c int, v float64) {
	this.Values[r][c] = v
}

func (this *MatrixFloat) Row() int {
	return len(this.Values)
}

func (this *MatrixFloat) Col() int {
	return len(this.Values[0])
}

func (this *MatrixFloat) AsVector() *Vector {
	var values []float64
	var size = this.Row()
	for i := 0; i < size; i++ {
		values = append(values, this.Values[i]...)
	}
	var ret = CreateVector2(values)
	return ret
}

func ListOfMatrixFloatAsVector(ms []*MatrixFloat) *Vector {
	var values []float64
	var size = len(ms)
	for i := 0; i < size; i++ {
		values = append(values, ms[i].AsVector().Values...)
	}
	var ret = CreateVector2(values)
	return ret
}

func CreateMatrixVector(rowsize int, colsize int, vectorsize int) *MatrixVector {
	var values [][]*Vector

	for i := 0; i < rowsize; i++ {
		var row []*Vector
		for j := 0; j < colsize; j++ {
			row = append(row, CreateVector(vectorsize))
		}

		values = append(values, row)
	}
	var ret = new(MatrixVector)
	ret.Values = values
	return ret
}

func (this *MatrixVector) Size() int {
	return len(this.Values) * len(this.Values[0])
}

func (this *MatrixVector) Get(r int, c int) *Vector {
	return this.Values[r][c]
}

func (this *MatrixVector) Put(r int, c int, v *Vector) {
	this.Values[r][c] = v
}

func (this *MatrixVector) Row() int {
	return len(this.Values)
}

func (this *MatrixVector) Col() int {
	return len(this.Values[0])
}
