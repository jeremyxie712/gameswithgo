package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]storyarc

type storyarc struct {
	title      string    `json:"title"`
	paragraphs []string  `json:"story"`
	Options    []options `json:"options"`
}

type options struct {
	text string `json:"text"`
	arc  string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
