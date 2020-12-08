package jsonparse

type Story map[string]storyarc

type storyarc struct {
	title   string    `json:"title"`
	story   []string  `json:"story"`
	Options []options `json:"options"`
}

type options struct {
	text string `json:"text"`
	arc  string `json:"arc"`
}
