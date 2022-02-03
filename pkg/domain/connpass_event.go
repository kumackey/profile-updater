package domain

import (
	"time"
)

type connpassEvent struct {
	title         string
	link          string
	ownerNickName string
	startedAt     time.Time
}

func (z *connpassEvent) toMarkdown(userNickName string) string {
	markdown := "- " + convertTimeToString(z.startedAt) + " "
	if z.isOwner(userNickName) {
		markdown += "**Organizer** "
	}

	markdown += "[" + z.title + "](" + z.link + ")"

	return markdown
}

func (z *connpassEvent) isOwner(userNickName string) bool {
	return z.ownerNickName == userNickName
}

func NewConpassEvent(title, link, ownerNickname string, startedAt time.Time) *connpassEvent {
	return &connpassEvent{title: title, link: link, ownerNickName: ownerNickname, startedAt: startedAt}
}
