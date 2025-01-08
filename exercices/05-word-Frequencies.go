package main

import ( 
	"fmt"
	"exercices/utils"
)

func main(){

	var strSlice []string
	var varInp string
	m := make(map[string]int)

	fmt.Println("Enter sentences or 'q' to stop")

	for {
		fmt.Scan(&varInp)
		if varInp == "q"{
			break
		}

		strSlice = append(strSlice, varInp)
	}

	next_ := utils.Iterator(strSlice) 
	for {
		str, finished := next_()
		if !finished {
			break
		}

		_, exists := m[str]

		if !exists {
			m[str] = 1
		} else {
			m[str]++
		}
	}

	fmt.Println(m)
}