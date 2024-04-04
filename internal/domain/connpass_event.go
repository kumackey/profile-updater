package domain

import (
	"time"
)

type ConnpassEvent struct {
	title         string
	link          string
	ownerNickname string
	startedAt     time.Time
}

func (z *ConnpassEvent) toMarkdown(userNickname string) string {
	markdown := "- " + convertTimeToString(z.startedAt) + " "
	if z.isOwner(userNickname) {
		markdown += "**Organizer** "
	}

	markdown += "[" + z.title + "](" + z.link + ")"

	return markdown
}

func (z *ConnpassEvent) isOwner(userNickname string) bool {
	return z.ownerNickname == userNickname
}

func NewConpassEvent(title, link, ownerNickname string, startedAt time.Time) *ConnpassEvent {
	return &ConnpassEvent{title: title, link: link, ownerNickname: ownerNickname, startedAt: startedAt}
}
