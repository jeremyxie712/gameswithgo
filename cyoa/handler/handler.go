package handler

import (
	"gameswithgo/cyoa"
	"gameswithgo/cyoa/info"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type handler struct {
	Information info.Information
}

func (h *handler) serveHTTP() {
	fileHandler := cyoa.JSONHandler{Information: h.Information}
	f, err := fileHandler.GetContent()
	if err != nil {
		log.Fatal(err)
	}
	urlHandler := h.MapHandler(f)
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(h.Information.GetPort(), nil)
}

func (h *handler) MapHandler(stories map[string]cyoa.Chapter) http.HandlerFunc {
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
