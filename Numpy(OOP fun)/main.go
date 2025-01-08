package main 

import (
	"fmt"
	"crazynumpy/numpy"
)

func main() {
	fmt.Println("Vectors:")
	v1 := numpy.NewVector(1.1, 2.0, 3.0)
	v2 := numpy.NewVector(4.0, 5.0, 6.0)
	fmt.Println(v1.String())
	fmt.Println(v2.String())
	v1.Add(v2)
	fmt.Println(v1.String())

	fmt.Println("Matrix:")
	m1 := numpy.Identity(4)
	m2 := numpy.Identity(4)
	fmt.Println(m1.String())
	fmt.Println(m2.String())
	m1.Mul(m2)
	fmt.Println(m1.String())
	m1.Add(m2)
	fmt.Println(m1.String())
	fmt.Println(m1.Rank())
	fmt.Println(m1.Det())
	m1.Transpose()
	fmt.Println(m1.String())
}

