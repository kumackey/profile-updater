package domain

import (
	"time"
)

type ZennArticle struct {
	Title       string
	Link        string
	EnClosure   EnClosure
	PublishedAt time.Time
}

type EnClosure struct {
	URL string
}

func (z *ZennArticle) toMarkdown() string {
	return "- [" + z.Title + "](" + z.Link + ")"
}
