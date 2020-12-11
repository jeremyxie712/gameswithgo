package process

import (
	"encoding/json"
	"gameswithgo/cyoa/info"
	"io/ioutil"
	"log"
)

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type JSONHandler struct {
	Information info.Information
}

func (fh *JSONHandler) GetContent() (map[string]Chapter, error) {
	Story := make(map[string]Chapter)
	f, err := ioutil.ReadFile(fh.Information.GetFilePath())
	if err != nil {
		log.Println("Cannot read the JSON file.")
		log.Fatal(err)
	}
	if err := json.Unmarshal(f, &Story); err != nil {
		return nil, err
	}
	return Story, nil
}
