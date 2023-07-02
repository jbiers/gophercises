package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type gameScore struct {
	correctAnswers int
	wrongAnswers   int
}

var score gameScore

func initializeScore(numProblems int) {
	score = gameScore{
		correctAnswers: 0,
		wrongAnswers:   numProblems,
	}
}

func loadProblems(filename string) [][]string {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error loading problem file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	problems, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading problem file:", err)
	}

	return problems
}

func questionUser(problems [][]string) {
	fmt.Printf("This quiz is %v questions long. Get ready.\n\n", len(problems))

	reader := bufio.NewReader(os.Stdin)

	for _, problem := range problems {
		fmt.Println(problem[0])

		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading answer:", err)
		}

		answer = answer[:len(answer)-1]

		if answer != problem[1] {
			fmt.Printf("Wrong. The answer is %v\n\n", problem[1])
		} else {
			score.wrongAnswers--
			score.correctAnswers++

			fmt.Printf("Correct.\n\n")
		}

	}
}

func showResults() {
	fmt.Printf("Results:\n")
	fmt.Printf("Wrong answers: %v\n", score.wrongAnswers)
	fmt.Printf("Correct answers: %v\n", score.correctAnswers)

}

func main() {
	// hardcoded filename?
	problems := loadProblems("problems.csv")

	initializeScore(len(problems))

	questionUser(problems)
	showResults()
}
