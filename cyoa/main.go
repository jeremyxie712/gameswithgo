package main

import (
	"flag"
	"gameswithgo/cyoa/info"
	"log"
	"strconv"
)

var (
	filename = flag.String("file", "/Users/JeremyXie/go/src/gameswithgo/cyoa/gopher.json", "the JSON file with CYOA story")
	port     = flag.Int("port", 8080, "the port start CYOA Web Application")
	tplPath  = flag.String("template", "/Users/JeremyXie/go/src/gameswithgo/cyoa/template/index.html", "the path to the template html file.")
)

func main() {
	flag.Parse()

	if len(*filename) == 0 || *filename == "" || *port == 0 {
		log.Println("Empty filepath or port, please check.")
	}
	config := info.Information{FilePath: *filename, LisPort: strconv.Itoa(*port), TemPath: *tplPath}
	hand := handler.pathHandler{config}
	hand.serveHTTP()
}
