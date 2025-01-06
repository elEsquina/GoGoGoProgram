package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var inp string
	var curr, best, limit, rang int


	best = loadBestScore()
	if best == -1 {
		rang, limit = selectDifficulty()
		best = game(rang, limit)
	}

	for {
		fmt.Println("Play again (y / yes)? Your best score is", best)
		fmt.Scan(&inp)

		if inp == "y" || inp == "yes" {
			rang, limit = selectDifficulty()
			curr = game(rang, limit)
			best = min(curr, best)

			saveBestScore(best)
		} else {
			break
		}
	}
}

func loadBestScore() int {
	file, err := os.Open("bestScore.txt")
	if err != nil {
		return -1
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil || count == 0 {
		return -1
	}

	score, err := strconv.Atoi(string(data[:count]))
	if err != nil {
		return -1
	}

	return score
}

func saveBestScore(score int) {
	file, err := os.Create("bestScore.txt")
	if err != nil {
		fmt.Println("Error saving score:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(strconv.Itoa(score)))
	if err != nil {
		fmt.Println("Error writing to file:", err)
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

func game(rang int, limit int) int {
	fmt.Println("Guess a number between 0 and", rang)

	random := rand.Intn(rang)
	var guess, count int

	for {
		if count >= limit {
			fmt.Println("You lose! The correct number was", random)
			break
		}

		fmt.Print("Enter your guess: ")
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
