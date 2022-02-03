package domain

import (
	"time"
)

type ZennArticle struct {
	title       string
	link        string
	publishedAt time.Time
}

func (z *ZennArticle) toMarkdown() string {
	publishedAt := convertTimeToString(z.publishedAt)

	return "- " + publishedAt + " [" + z.title + "](" + z.link + ")"
}

func NewZennArticle(title, link string, publishedAt time.Time) *ZennArticle {
	return &ZennArticle{title: title, link: link, publishedAt: publishedAt}
}
