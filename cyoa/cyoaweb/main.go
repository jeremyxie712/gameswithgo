package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the file with the story.")
	flag.Parse()
	fmt.Printf("Using the story %v\n", *filename)

	f, err := os.Open(*filenname)
	if err != nil {
		log.Fatal(err)
	}

	d := json.NewDecoder(f)
	var story jsonparse.Story

}
