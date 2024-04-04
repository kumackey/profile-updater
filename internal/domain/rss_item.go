package domain

import (
	"context"
	"errors"
	"net/url"
	"time"
)

type RssItem struct {
	title       string
	link        string
	publishedAt time.Time
}

func (r RssItem) ToMarkdown() string {
	publishedAt := convertTimeToString(r.publishedAt)

	return "- " + publishedAt + " [" + r.title + "](" + r.link + ")"
}

func (r RssItem) SortOrder() int64 {
	return r.publishedAt.Unix()
}

func NewRssItem(title, link string, publishedAt time.Time) RssItem {
	return RssItem{title: title, link: link, publishedAt: publishedAt}
}

var (
	ErrRssNotFound            = errors.New("rss not found")
	ErrRssInternalServerError = errors.New("rss internal server error")
	ErrRssUnknownError        = errors.New("rss unknown error")
)

type RssClient interface {
	FetchItems(ctx context.Context, url *url.URL) ([]RssItem, error)
}
