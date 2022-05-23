package domain

import (
	"time"
)

type qiitaArticle struct {
	title       string
	link        string
	lgtms       int
	publishedAt time.Time
}

func (z *qiitaArticle) toMarkdown() string {
	publishedAt := convertTimeToString(z.publishedAt)

	return "- " + publishedAt + " [" + z.title + "](" + z.link + ")"
}

func NewQiitaArticle(title, link string, lgtms int, publishedAt time.Time) *qiitaArticle {
	return &qiitaArticle{title: title, link: link, lgtms: lgtms, publishedAt: publishedAt}
}
