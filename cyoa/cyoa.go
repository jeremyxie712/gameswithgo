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
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <div class="story-content">
        {{range .Story}}
            <p>{{.}}</p>
        {{end}}
    </div>
    <div class="variations-links">
        {{if .Options }}
            <ul>
                {{range .Options}}
                    <li class="link-container">
                        <a href="{{.Arc}}">{{.Text}}</a>
                    </li>
                {{end}}
            </ul>
        {{else}}
            <p ><b>The End !</b></p>
        {{end}}
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
