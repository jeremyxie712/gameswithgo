package main

import (
	"flag"
	"fmt"
	"gameswithgo/cyoa"
	"log"
	"net/http"
	"os"
)

const (
	templateForStory = "/Users/JeremyXie/go/src/gameswithgo/cyoa/cmd/index.html"
)

var (
	filename = flag.String("file", "/Users/JeremyXie/go/src/gameswithgo/cyoa/gopher.json", "the JSON file with CYOA story")
	port     = flag.Int("port", 8080, "the port to start CYOA Web Application")
)

func main() {
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.ParseJSON(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
