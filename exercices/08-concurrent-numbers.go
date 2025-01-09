package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	var slice []int

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		waitGroup.Add(1)
		go squareNumber(i)
	}

	waitGroup.Wait()
	fmt.Println("done")
}

func squareNumber(n int) {
	defer waitGroup.Done()
	fmt.Println("square of", n, "is", n*n)
}