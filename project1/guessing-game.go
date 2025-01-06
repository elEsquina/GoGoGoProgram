package main

import (
    "fmt"
    "math/rand/v2"
	"os"
)

func main(){

	var inp string
	var curr, best, limit, rang int

	rang, limit = selectDifficulty()
	best = game(rang, limit)

	for {

		fmt.Println("Play again (y / yes) ? Your best score is", best)
		fmt.Scan(&inp)

		if inp == "y" || inp == "yes" {
			rang, limit = selectDifficulty()
			curr = game(rang, limit)
			best = min(curr, best)	
		}
		else {
			break
		}

	}
}

func selectDifficulty() (int, int) {
	var diff string

	fmt.Println("Choose your difficulty: easy, medium, hard")
	fmt.Scan(&diff)

	switch diff {
		case "easy":
			return 10, 5
		case "medium":
			return 50, 10
		case "hard":
			return 100, 15
		default:
			fmt.Println("Invalid difficulty")
			os.Exit(1)
			return 0, 0
	}
}


func game(rang int, limit int) int{

	fmt.Println("Guess a number between 0 and", rang)

	var random int = rand.IntN(rang)
	var guess, count int

	for { 

		if count > limit {
			fmt.Println("Looser")	
			break
		}

		fmt.Scan(&guess)
		count++

		if guess < random {
			fmt.Println("Too low!")
		} else if guess > random {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You got it!")
			break
		}

	}

	return count
}