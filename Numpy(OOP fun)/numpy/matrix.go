package numpy

import (
	"errors"
	"fmt"
	"math"
)

var eps float64 = 1e-9

type Matrix struct {
	Coords [][]float64
	Rows int
	Cols int
}

func Zeros(rows, cols int) Matrix {
	coords := make([][]float64, rows)
	for i := range coords {
		coords[i] = make([]float64, cols)
	}
	return Matrix{coords: coords, rows: rows, cols: cols}
}

func Ones(rows, cols int) Matrix {
	coords := make([][]float64, rows)
	for i := range coords {
		coords[i] = make([]float64, cols)
		for j := range coords[i] {
			coords[i][j] = 1.0
		}
	}
	return Matrix{coords: coords, rows: rows, cols: cols}
}

func Identity(size int) Matrix {
	coords := make([][]float64, size)
	for i := range coords {
		coords[i] = make([]float64, size)
		coords[i][i] = 1.0
	}
	return Matrix{coords: coords, rows: size, cols: size}
}

func (m *Matrix) Copy() Matrix {
	var ret Matrix = Zeros(m.rows, m.cols)

	for i := 0; i < m.rows; i++{
		for j := 0; j < m.cols; j++{
			ret.coords[i][j] = m.coords[i][j]
		}
	}

	return ret
}

func NewMatrix(coords [][]float64) (Matrix, error) {
	rows := len(coords)
	if rows == 0 {
		var z Matrix 
		return z, errors.New("matrix must have at least one row")
	}
	cols := len(coords[0])
	for i := 1; i < rows; i++ {
		if len(coords[i]) != cols {
			var z Matrix 
			return z, errors.New("inconsistent row lengths")
		}
	}
	return Matrix{coords: coords, rows: rows, cols: cols}, nil
}

func (m *Matrix) Add(other *Matrix) error {
	if m.rows != other.rows || m.cols != other.cols {
		return errors.New("dimension mismatch")
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.coords[i][j] += other.coords[i][j]
		}
	}
	return nil
}

func (m *Matrix) Sub(other *Matrix) error {
	if m.rows != other.rows || m.cols != other.cols {
		return errors.New("dimension mismatch")
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.coords[i][j] -= other.coords[i][j]
		}
	}
	return nil
}

func (m *Matrix) Mul(other *Matrix) error {
	if m.cols != other.rows {
		return errors.New("dimension mismatch")
	}
	result := Zeros(m.rows, other.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			for k := 0; k < m.cols; k++ {
				result.coords[i][j] += m.coords[i][k] * other.coords[k][j]
			}
		}
	}
	m.coords = result.coords
	m.cols = result.cols
	return nil
}

func (m *Matrix) Scale(scalar float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.coords[i][j] *= scalar
		}
	}
}

func (m *Matrix) Transpose() {
	transposed := make([][]float64, m.cols)
	for i := range transposed {
		transposed[i] = make([]float64, m.rows)
		for j := range transposed[i] {
			transposed[i][j] = m.coords[j][i]
		}
	}
	m.coords = transposed
	m.rows, m.cols = m.cols, m.rows
}

func (m *Matrix) Det() (float64, error) {
	if m.rows != m.cols {
		return 0, errors.New("matrix is not square")
	}

	n := m.rows
	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
		copy(a[i], m.coords[i])
	}

	det := 1.0
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if math.Abs(a[j][i]) > math.Abs(a[pivot][i]) {
				pivot = j
			}
		}
		if math.Abs(a[pivot][i]) < eps {
			return 0, nil
		}

		a[i], a[pivot] = a[pivot], a[i]
		if i != pivot {
			det = -det
		}
		det *= a[i][i]

		for j := i + 1; j < n; j++ {
			a[i][j] /= a[i][i]
		}
		for j := 0; j < n; j++ {
			if j != i {
				for k := i + 1; k < n; k++ {
					a[j][k] -= a[j][i] * a[i][k]
				}
			}
		}
	}
	return det, nil
}


func (m *Matrix) Rank() int {
	n := m.rows
	m.Transpose()
	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
		copy(a[i], m.coords[i])
	}
	m.Transpose()

	rank := 0
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if math.Abs(a[j][i]) > math.Abs(a[pivot][i]) {
				pivot = j
			}
		}
		if math.Abs(a[pivot][i]) > eps {
			rank++
			a[i], a[pivot] = a[pivot], a[i]
			for j := i + 1; j < n; j++ {
				a[i][j] /= a[i][i]
			}
			for j := 0; j < n; j++ {
				if j != i {
					for k := i + 1; k < n; k++ {
						a[j][k] -= a[j][i] * a[i][k]
					}
				}
			}
		}
	}
	return rank
}


/*
func (m *Matrix) Inverse() error {
	if m.rows != m.cols {
		return Error.New("Matrix is not square")
	}
	if m.Det() == 0 {
		return Error.New("Matrix is singular")
	}

	var n int = m.rows
	var _ Matrix = Identity(n)

	// TODO: Implement Gaussian elimination
}
*/

func (m *Matrix) String() string {
	str := ""
	for i := 0; i < m.rows; i++ {
		str += fmt.Sprintf("%v\n", m.coords[i])
	}
	return str
}