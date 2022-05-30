package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.Question)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.Answer {
				fmt.Println("Correct!")
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// Problem is a struct that represents a pair of question answer from the CSV file
type Problem struct {
	Question string
	Answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
