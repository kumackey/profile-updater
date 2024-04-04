package domain

import "sort"

type Item interface {
	ToMarkdown() string
	SortOrder() int64
}

func ToMarkdown[T Item](items []T, limit int) string {
	md := "\n"

	sort.Slice(items, func(i, j int) bool {
		return items[j].SortOrder() < items[i].SortOrder()
	})

	if len(items) > limit {
		items = items[:limit]
	}

	for _, item := range items {
		md = md + item.ToMarkdown() + "\n"
	}

	return md
}
