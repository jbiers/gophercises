package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type gameScore struct {
	correctAnswers  int
	wrongAnswers    int
	remainingAnwers int
}

var score gameScore

var reader *bufio.Reader

var timeTerminated chan int
var questionsTerminated chan int

func initializeScore(numProblems int) {
	score = gameScore{
		correctAnswers:  0,
		wrongAnswers:    0,
		remainingAnwers: numProblems,
	}
}

func loadProblems(filename string) [][]string {
	file, err := os.Open(filename)
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

func questionsLoop(problems [][]string) {
	fmt.Printf("This quiz is %v questions long. Get ready.\n\n", len(problems))

	reader = bufio.NewReader(os.Stdin)

	for _, problem := range problems {
		question, result := problem[0], problem[1]

		askQuestion(question, result)

	}

	questionsTerminated <- 1
}

func askQuestion(q string, a string) {
	fmt.Println(q)

	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading answer:", err)
	}

	answer = answer[:len(answer)-1]

	if answer != a {
		score.wrongAnswers++
		score.remainingAnwers--

		fmt.Printf("Wrong. The answer is %v\n\n", a)
	} else {
		score.correctAnswers++
		score.remainingAnwers--
		fmt.Printf("Correct.\n\n")
	}
}

func countTime(s time.Duration) {
	startTime := time.Now()

	for {
		if time.Since(startTime) > time.Second*s {
			break
		}
	}

	timeTerminated <- 1
}

func showResults() {
	fmt.Printf("Results:\n")
	fmt.Printf("Wrong answers: %v\n", score.wrongAnswers)
	fmt.Printf("Correct answers: %v\n", score.correctAnswers)
	fmt.Printf("Remaining answers: %v\n", score.remainingAnwers)
}

func main() {
	filename := flag.String("file", "problems.csv", "name of the problem set file")
	timeLimit := flag.Int("time", 30, "time limit to answer the questions, in seconds")

	flag.Parse()

	problems := loadProblems(*filename)
	initializeScore(len(problems))

	timeTerminated, questionsTerminated = make(chan int), make(chan int)

	go questionsLoop(problems)
	go countTime(time.Duration(*timeLimit))

	select {
	case _ = <-timeTerminated:
		fmt.Printf("\nSorry, you ran out of time.\n\n")
	case _ = <-questionsTerminated:
		fmt.Printf("\nYou finished the quiz\n\n")
	}

	showResults()
}
