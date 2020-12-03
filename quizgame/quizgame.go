package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	problemsFile string
	timeInterval int
	numCorrect   int
	comparison   bool
	userAnswer   string
)

type problems struct { //This way, the type problems could be used by other functions as well
	q string
	a string
}

func preprocess(input string) string {
	return strings.ToLower(strings.Trim(input, "\n\t\r\n"))
}

func parseProblems(lines [][]string) []problems {
	if len(lines) == 0 || len(lines[0]) == 0 {
		fmt.Print("Failed to open csv files at parsing.")
		os.Exit(1)
	}
	res := make([]problems, len(lines))
	for i, line := range lines {
		res[i] = problems{
			q: line[0],
			a: line[1],
		}
	}
	return res
}
func parseLines(lines [][]string) [][]string {
	if len(lines) == 0 || len(lines[0]) == 0 {
		fmt.Print("Failed to open csv filles at parsing.")
		os.Exit(1)
	}
	res := make([][]string, len(lines))
	for i := range res {
		res[i] = make([]string, len(lines[0]))
	}
	for i, line := range lines {
		res[i][0] = line[0]
		res[i][1] = strings.TrimSpace(line[1])
	}
	return res
}
func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'questions,answers")
	// flag.StringVar(&problemsFile, "probs", "problems.csv", "This is the problem set.")
	// flag.IntVar(&timeInterval, "time interval", 30, "Time for a single question.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Failed to open the csv file", err)
		os.Exit(1)
	}
	defer file.Close()

	lineInfo, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln("Failed to open the problem set", err)
		os.Exit(1)
	}

	parsedLines := parseProblems(lineInfo)
	parsedSlices := parseLines(lineInfo)
	if comparison {
		fmt.Println("Parsing using problem struct: ")
		fmt.Println(parsedLines[0])
		fmt.Println("Parsing using two dimensional slices: ")
		fmt.Println(parsedSlices[0])
	}

	for idx, p := range parsedSlices {
		fmt.Printf("Question: #%d: What is %s = ?\n", idx+1, p[0])
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == p[1] {
			numCorrect++
		} else {
			fmt.Println("Oops! Not right. Keep it up!")
		}
	}
	fmt.Printf("You have scored %s questions out of %s\n", numCorrect, len(parsedSlices))
}
