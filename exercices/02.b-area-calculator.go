package main 

import "fmt"

func main(){
	var length, width float64
	length = 1.5
	width = 2.5
	fmt.Println("Area:", area(length, width))
}

func area(length, width float64) float64{
	return length * width * 3.28 *  3.28
}