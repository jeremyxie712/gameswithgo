package main

import (
	"log"
	"os"
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "gopher.json", "the file with the story.")
	flag.Parse()
	fmt.Printf("Using the story %v\n", *filename)

	f, err := os.Open(*filenname)
	if err != nil {[
		log.Fatal(err)
	]}
}
