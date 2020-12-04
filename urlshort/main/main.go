package main

import (
	"bytes"
	"flag"
	"fmt"
	"gameswithgo/urlshort"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
	// "github.com/gophercises/urlshort"
)

var (
	approvalToDB = flag.Bool("createDB", true, "Boolean value to determine whether to create DB or not.")
	pathToFile   = flag.String("paths", "paths.yml", "The path to the YAML file.")
)

func createDB() (*bolt.DB, error) {
	db, err := bolt.Open(*pathToFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())
	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("paths"))
		if err != nil {
			return err
		}
		if err := b.Put([]byte("/urlshort"), []byte("https://github.com/gophercises/urlshort")); err != nil {
			return err
		}
		if err := b.Put([]byte("/urlshort-final"), []byte("https://github.com/gophercises/urlshort/tree/solution")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return db, err
}

func turnFileToBytes(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open the file %v\n", filename)
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatalf("Cannot read the file %v\n", filename)
	}
	return buf.Bytes()
}

func main() {
	mux := defaultMux()

	flag.Parse()

	var handler http.Handler
	var err error
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	switch filepath.Ext(*pathToFile) {
	case "*.yml":
		handler, err := urlshort.YAMLHandler(turnFileToBytes(*pathToFile), mux)
		if err != nil {
			panic(err)
		}
	case "*.json":
		handler, err := urlshort.JSONHandler(turnFileToBytes(*pathToFile), mux)
		if err != nil {
			panic(err)
		}
	case "*.db":
		if *approvalToDB {
			db, err := createDB()
			if err != nil {
				panic(err)
			}
			handler = urlshort.BoltDBHandler(db, mux)
		}
	default:
		log.Fatalln("Paths file is unformatted.")
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
