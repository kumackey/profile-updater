package domain

import (
	"time"
)

type ZennArticle struct {
	title       string
	link        string
	publishedAt time.Time
}

func (z ZennArticle) ToMarkdown() string {
	publishedAt := convertTimeToString(z.publishedAt)

	return "- " + publishedAt + " [" + z.title + "](" + z.link + ")"
}

func (z ZennArticle) SortOrder() int64 {
	return z.publishedAt.Unix()
}

func NewZennArticle(title, link string, publishedAt time.Time) ZennArticle {
	return ZennArticle{title: title, link: link, publishedAt: publishedAt}
}
