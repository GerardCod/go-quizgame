package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

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

	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.Answer {
			fmt.Println("Correct!")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
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
