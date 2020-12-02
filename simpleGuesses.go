package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input a number between 1 and 100")
	scanner.Scan()

	low, high, count := 1 ,100, 0
	for {
		guess := low + (high - low) / 2
		fmt.Println("Guessed number:",guess)
		fmt.Println("(a) Too high")
		fmt.Println("(b) Too low")
		fmt.Println("(c) Correct")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			count++
		}else if response == "b" {
			low = guess + 1
			count++
		}else if response == "c" {
			fmt.Printf("correct! %d attemps in total", count)
			break
		}else{
			fmt.Println("invalid input")
		}

	}
}
