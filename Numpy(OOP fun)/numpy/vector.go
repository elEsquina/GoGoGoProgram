package numpy

import (
	"errors"
	"fmt"
)

type Vector struct {
	Coords []float64
	dim int
}

func NewVector(coords ...float64) Vector {
	return Vector{
		coords: coords,
		dim:    len(coords),
	}
}

func (v *Vector) Copy() Vector{
	return NewVector(v.coords...) 
}

func (v *Vector) Add(other *Vector) error {
	if v.dim != other.dim {
		return errors.New("dimension mismatch")
	}

	for i := 0; i < v.dim; i++ {
		v.coords[i] += other.coords[i]
	}
	return nil
}

func (v *Vector) Sub(other *Vector) error {
	if v.dim != other.dim {
		return errors.New("dimension mismatch")
	}

	for i := 0; i < v.dim; i++ {
		v.coords[i] -= other.coords[i]
	}
	return nil
}

func (v *Vector) Dot(other *Vector) (float64, error) {
	if v.dim != other.dim {
		return 0, errors.New("dimension mismatch")
	}

	result := 0.0
	for i := 0; i < v.dim; i++ {
		result += v.coords[i] * other.coords[i]
	}

	return result, nil
}

func (v *Vector) Scale(scalar float64) {
	for i := 0; i < v.dim; i++ {
		v.coords[i] *= scalar
	}
}

func (v *Vector) String() string {
	return fmt.Sprintf("%v", v.coords)
}

func (v *Vector) Dim() int {
	return v.dim
}
