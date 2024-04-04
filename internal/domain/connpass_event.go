package domain

import (
	"time"
)

type ConnpassEvent struct {
	title     string
	link      string
	startedAt time.Time
}

func (z *ConnpassEvent) ToMarkdown() string {
	return "- " + convertTimeToString(z.startedAt) + " " + "[" + z.title + "](" + z.link + ")"
}

func (z *ConnpassEvent) SortOrder() int64 {
	return z.startedAt.Unix()
}

func NewConpassEvent(title, link string, startedAt time.Time) *ConnpassEvent {
	return &ConnpassEvent{title: title, link: link, startedAt: startedAt}
}
