package domain

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToMarkdown_Connpass(t *testing.T) {
	tests := map[string]struct {
		events []ConnpassEvent
		limit  int
		output string
	}{
		"マークダウンに変換できること": {
			events: []ConnpassEvent{
				connpass("イベントの例1", "https://example.com/1", "2022-02-01T14:59:00+00:00"),
				connpass("イベントの例2", "https://example.com/2", "2022-02-01T15:00:00+00:00"),
			},
			limit:  5,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n- Feb 1 [イベントの例1](https://example.com/1)\n",
		},
		"イベント数を制限できること": {
			events: []ConnpassEvent{
				connpass("イベントの例1", "https://example.com/1", "2022-02-01T14:59:00+00:00"),
				connpass("イベントの例2", "https://example.com/2", "2022-02-01T15:00:00+00:00"),
			},
			limit:  1,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sort.Slice(test.events, func(i, j int) bool {
				return test.events[i].SortOrder() > test.events[j].SortOrder()
			})
			markdown := ToMarkdown(test.events, test.limit)
			assert.Equal(t, test.output, markdown)
		})
	}
}

func connpass(title, link, startedAt string) ConnpassEvent {
	sa, err := time.Parse(time.RFC3339, startedAt)
	if err != nil {
		panic("connpass, time.Parse failed")
	}
	return ConnpassEvent{title: title, link: link, startedAt: sa}
}

func TestToMarkdown_Qiita(t *testing.T) {
	tests := map[string]struct {
		input  []QiitaArticle
		output string
	}{
		"マークダウンに変換できること": {
			input: []QiitaArticle{
				qiita("記事の例1", "https://example.com/1", 5, "2022-02-01T14:59:00+00:00"),
				qiita("記事の例2", "https://example.com/2", 0, "2022-02-01T15:00:00+00:00"),
			},

			output: "\n- Feb 2 [記事の例2](https://example.com/2)\n- Feb 1, **5 LGTM** [記事の例1](https://example.com/1)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sort.Slice(test.input, func(i, j int) bool {
				return test.input[i].SortOrder() > test.input[j].SortOrder()
			})
			markdown := ToMarkdown(test.input, 5)
			assert.Equal(t, test.output, markdown)
		})
	}
}

func qiita(title, link string, lgtms int, publishedAt string) QiitaArticle {
	pa, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		panic("qiita, time.Parse failed")
	}
	return QiitaArticle{title: title, link: link, lgtms: lgtms, publishedAt: pa}
}

func TestToMarkdown_Zenn(t *testing.T) {
	tests := map[string]struct {
		input  []RssItem
		output string
	}{
		"マークダウンに変換できる": {
			input: []RssItem{
				item("記事の例1", "https://example.com/1", "2022-02-01T14:59:00+00:00"),
				item("記事の例2", "https://example.com/2", "2022-02-01T15:00:00+00:00"),
			},

			output: "\n- Feb 2 [記事の例2](https://example.com/2)\n- Feb 1 [記事の例1](https://example.com/1)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sort.Slice(test.input, func(i, j int) bool {
				return test.input[i].SortOrder() > test.input[j].SortOrder()
			})
			markdown := ToMarkdown(test.input, 5)
			assert.Equal(t, test.output, markdown)
		})
	}
}

func item(title, link, publishedAt string) RssItem {
	pa, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		panic("rss, time.Parse failed")
	}
	return RssItem{title: title, link: link, publishedAt: pa}
}
