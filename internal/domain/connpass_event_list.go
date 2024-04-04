package domain

import "sort"

type ConpassEventList []*ConnpassEvent

func (l ConpassEventList) ToProfileMarkdown(userNickname string) string {
	profileMarkdown := "\n"
	for _, event := range l {
		profileMarkdown = profileMarkdown + event.toMarkdown(userNickname) + "\n"
	}

	return profileMarkdown
}

func (l ConpassEventList) SortByPublishedAt() ConpassEventList {
	sort.Slice(l, func(i, j int) bool {
		// 開始時刻が遅い順
		return l[j].startedAt.Unix() < l[i].startedAt.Unix()
	})

	return l
}

func (l ConpassEventList) Limit(limit int) ConpassEventList {
	list := ConpassEventList{}
	count := 0
	for _, article := range l {
		if limit <= count {
			break
		}

		list = append(list, article)
		count += 1
	}

	return list
}
