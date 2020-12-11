package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type handler struct {
	s Story
	t *template.Template
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<head>
    <title>{{.Title}}</title>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-rc.2/css/materialize.min.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
  </head>
  <body class="grey lighten-4">
    <div class="row">
        <div class="col m12">
          <div class="card large horizontal z-depth-4 brown white-text">
            <div class="card-stacked">
              <h2 class="card-title">{{.Title}}</h2>
              <div class="card-content">
                {{range .Story}}
                  <p>{{.}}</p>
                {{end}}
              </div>
              <div class="card-action">
                {{range .Options}}
                  <a href="{{.Arc}}">{{.Text}}</p>
                {{end}}
              </div>
            </div>
          </div>
        </div>
      </div>
  </body>`

func ParseJSON(f io.Reader) (Story, error) {
	dec := json.NewDecoder(f)
	story := make(Story)
	if err := dec.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story, tmpl *template.Template) http.Handler {
	if tmpl == nil {
		tmpl = tpl
	}
	return handler{s, tpl}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" { //This is to prevent goes to localhost port but no path
		path = "/intro"
	}
	path = path[1:] // "/intro" --> "intro"
	//   ["intro"]
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went Wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter Not Found.", http.StatusNotFound)

}
