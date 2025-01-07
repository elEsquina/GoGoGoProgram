package main

import "fmt"
import "exercices/utils"

func main() {
	var numbers []int


	numbers = append(numbers, 10)
	numbers = append(numbers, 3)
	numbers = append(numbers, 5)
	numbers = append(numbers, 2)
	numbers = append(numbers, 5)
	numbers = append(numbers, -1)

	next_ := utils.Iterator(numbers) 
	for {
		num, err := next_()
		if !err {
			break
		}
		fmt.Println(num)
	}

}

