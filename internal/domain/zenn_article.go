package domain

import (
	"time"
)

type zennArticle struct {
	title       string
	link        string
	publishedAt time.Time
}

func (z *zennArticle) toMarkdown() string {
	publishedAt := convertTimeToString(z.publishedAt)

	return "- " + publishedAt + " [" + z.title + "](" + z.link + ")"
}

func NewZennArticle(title, link string, publishedAt time.Time) *zennArticle {
	return &zennArticle{title: title, link: link, publishedAt: publishedAt}
}
