package main

import (
	"net/http"
)

const message = "Halo Munchen"
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(message))
	})

	_ = http.ListenAndServe(":8080",mux)
}
