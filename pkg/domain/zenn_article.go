package domain

import (
	"sort"
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

func (z *ZennArticle) ToMarkdown() string {
	return "[" + z.Title + "](" + z.Link + ")"
}

type ZennArticles []*ZennArticle

func (z ZennArticles) SortByPublishedAt() ZennArticles {
	sort.Slice(z, func(i, j int) bool {
		// 公開が遅い順
		return z[j].PublishedAt.Unix() < z[i].PublishedAt.Unix()
	})

	return z
}
