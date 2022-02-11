package domain

import (
	"time"
)

type qiitaArticle struct {
	title       string
	link        string
	publishedAt time.Time
}

func (z *qiitaArticle) toMarkdown() string {
	publishedAt := convertTimeToString(z.publishedAt)

	return "- " + publishedAt + " [" + z.title + "](" + z.link + ")"
}

func NewQiitaArticle(title, link string, publishedAt time.Time) *qiitaArticle {
	return &qiitaArticle{title: title, link: link, publishedAt: publishedAt}
}
