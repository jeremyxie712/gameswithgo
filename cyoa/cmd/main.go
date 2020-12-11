package main

import (
	"flag"
	"gameswithgo/cyoa/info"
	"log"
	"strconv"
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

	if len(*filename) == 0 || *filename == "" || *port == 0 {
		log.Println("Empty filepath or port, please check.")
	}
	config := info.Information{FilePath: *filename, LisPort: strconv.Itoa(*port), TemPath: templateForStory}
	handler := path_handler.pathHandler{config}
	handler.serveHTTP()

}
