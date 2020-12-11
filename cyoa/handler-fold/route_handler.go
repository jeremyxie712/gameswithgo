package handler

import (
	"gameswithgo/cyoa/info"
	"gameswithgo/cyoa/process"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type pathHand struct {
	Information info.Information
}

func (h *pathHand) serveHTTP() {
	fileHandler := process.JSONHandler{Information: h.Information}
	f, err := fileHandler.GetContent()
	if err != nil {
		log.Fatal(err)
	}
	urlHandler := h.MapHandler(f)
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(":"+h.Information.GetPort(), nil)
}

func (h *pathHand) MapHandler(stories map[string]process.Chapter) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimSpace(r.URL.Path)
		t := template.Must(template.ParseFiles(h.Information.GetTmplPath()))
		if story, ok := stories[url]; ok {
			t.Execute(w, story)
			return
		}
		t.Execute(w, stories["intro"])
	})
}
