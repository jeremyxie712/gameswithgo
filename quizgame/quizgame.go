package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	filename     = flag.String("path", "problems.csv", "Path to csv files containing 'question-answer' pairs")
	timeInterval = flag.Int("limit", 30, "Time allowed for each question in seconds.")
	problemsFile string
	numCorrect   int
	numAnswered  int
	comparison   bool
	need         bool
	userAnswer   string
)

type problems struct { //This way, the type problems could be used by other functions as well
	q string
	a string
}

func loadCSV(path string) [][]string {
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
	return lineInfo
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

func sanityCheck(comparison bool, parsedLines []problems, parsedSlices [][]string) {
	if comparison {
		fmt.Println("Parsing using problem struct: ")
		fmt.Println(parsedLines[0])
		fmt.Println("Parsing using two dimensional slices: ")
		fmt.Println(parsedSlices[0])
	}
}

func quizTime(parsedQuestions [][]string) {
	if len(parsedQuestions) == 0 || len(parsedQuestions[0]) == 0 || parsedQuestions == nil {
		fmt.Println("Failed to read parsed csv.")
		os.Exit(1)
	}
	timer := time.NewTimer(time.Duration(*timeInterval) * time.Second)
	for idx, p := range parsedQuestions {
		fmt.Printf("Question #%v: What is %v = ", idx+1, p[0])
		answerChan := make(chan string)
		go func() {
			scanner.Scan()
			userAnswer := scanner.Text()
			answerChan <- userAnswer
		}()
		select {
		case <-timer.C:
			report(parsedQuestions)
			return
		case answer := <-answerChan:
			if answer == p[1] {
				numCorrect++
			} else {
				fmt.Println("Oops. Almost there, keep it up!")
			}
			numAnswered++
		}
	}
}

func report(questions [][]string) {
	fmt.Printf("\nYou have answered %d questions, total questions: %d, correct number of questions: %d", numAnswered, len(questions), numCorrect)

}

func main() {
	lines := loadCSV(*filename)

	parsedLines := parseProblems(lines)
	parsedSlices := parseLines(lines)

	if need {
		sanityCheck(comparison, parsedLines, parsedSlices)
	}

	quizTime(parsedSlices)
}
