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
)

type problems struct { //This way, the type problems could be used by other functions as well
	q string
	a string
}

func preprocess(input string) string {
	return strings.ToLower(strings.Trim(input, ""))
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
	fmt.Println(parsedLines[0])
}
