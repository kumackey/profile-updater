package domain

import (
	"strconv"
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

	text := "- " + publishedAt

	if z.lgtms != 0 {
		text += ", **" + strconv.Itoa(z.lgtms) + " LGTM**"
	}

	text += " [" + z.title + "](" + z.link + ")"

	return text
}

func NewQiitaArticle(title, link string, lgtms int, publishedAt time.Time) *qiitaArticle {
	return &qiitaArticle{title: title, link: link, lgtms: lgtms, publishedAt: publishedAt}
}
