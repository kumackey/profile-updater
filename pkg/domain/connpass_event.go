package domain

import (
	"time"
)

type connpassEvent struct {
	title         string
	link          string
	ownerNickname string
	startedAt     time.Time
}

func (z *connpassEvent) toMarkdown(userNickname string) string {
	markdown := "- " + convertTimeToString(z.startedAt) + " "
	if z.isOwner(userNickname) {
		markdown += "**Organizer** "
	}

	markdown += "[" + z.title + "](" + z.link + ")"

	return markdown
}

func (z *connpassEvent) isOwner(userNickname string) bool {
	return z.ownerNickname == userNickname
}

func NewConpassEvent(title, link, ownerNickname string, startedAt time.Time) *connpassEvent {
	return &connpassEvent{title: title, link: link, ownerNickname: ownerNickname, startedAt: startedAt}
}
