package main 

import "fmt"

func main(){
	
	var grades []float64
	var grade float64
	var total float64
	var average float64

	for {
		fmt.Println("Enter a grade or type 'quit' to finish:")
		_, err := fmt.Scan(&grade)
		if err != nil {
			break
		}
		grades = append(grades, grade)
		total += grade
	}

	if len(grades) == 0 {
		fmt.Println("No grades were entered.")
	}

	average = total /  float64(len(grades))

	fmt.Println("Average:", average)

	if average >= 90 {
		fmt.Println("A")
	} else if average >= 80 {
		fmt.Println("B")
	} else if average >= 70 {
		fmt.Println("C")
	} else if average >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}



}