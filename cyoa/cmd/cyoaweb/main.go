package main

import (
	"flag"
	"fmt"
	"gameswithgo/cyoa"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "/Users/JeremyXie/go/src/gameswithgo/cyoa/gopher.json", "the file with the story.")
	flag.Parse()
	fmt.Printf("Using the story %v.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", story)

}
