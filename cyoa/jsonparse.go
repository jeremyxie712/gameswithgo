package jsonparse

type storyarc struct {
	title   string
	story   []string
	Options []options
}

type options struct {
	text string
	arc  string
}
