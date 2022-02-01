package domain

import (
	"time"
)

const (
	JSTOffset = 9 * 60 * 60
	asiaTokyo = "Asia/Tokyo"
)

var locationJST = time.FixedZone(asiaTokyo, JSTOffset)

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
	publishedAt := z.PublishedAt.In(locationJST).Format("Jan 2, ")

	return "- " + publishedAt + "[" + z.Title + "](" + z.Link + ")"
}
