package domain

type Item interface {
	ToMarkdown() string
	SortOrder() int64
}
