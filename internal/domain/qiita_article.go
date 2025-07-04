package domain

import (
	"strconv"
	"time"
)

type QiitaArticle struct {
	title       string
	link        string
	lgtms       int
	publishedAt time.Time
}

func (q QiitaArticle) LGTMs() int {
	return q.lgtms
}

func (q QiitaArticle) ToMarkdown() string {
	publishedAt := convertTimeToString(q.publishedAt)

	text := "- " + publishedAt

	if q.lgtms != 0 {
		text += ", **" + strconv.Itoa(q.lgtms) + " LGTM**"
	}

	text += " [" + q.title + "](" + q.link + ")"

	return text
}

func (q QiitaArticle) SortOrder() int64 {
	return q.publishedAt.Unix()
}

func NewQiitaArticle(title, link string, lgtms int, publishedAt time.Time) QiitaArticle {
	return QiitaArticle{title: title, link: link, lgtms: lgtms, publishedAt: publishedAt}
}
