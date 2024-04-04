package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConpassEventList_ToProfileMarkdown(t *testing.T) {
	publishedAt1, _ := time.Parse(time.RFC3339, "2022-02-01T14:59:00+00:00")
	publishedAt2, _ := time.Parse(time.RFC3339, "2022-02-01T15:00:00+00:00")

	tests := map[string]struct {
		input  []ConnpassEvent
		limit  int
		output string
	}{
		"マークダウンに変換できる": {
			input: []ConnpassEvent{
				{title: "イベントの例1", link: "https://example.com/1", startedAt: publishedAt1},
				{title: "イベントの例2", link: "https://example.com/2", startedAt: publishedAt2},
			},
			limit:  5,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n- Feb 1 [イベントの例1](https://example.com/1)\n",
		},
		"イベント数を制限": {
			input: []ConnpassEvent{
				{title: "イベントの例1", link: "https://example.com/1", startedAt: publishedAt1},
				{title: "イベントの例2", link: "https://example.com/2", startedAt: publishedAt2},
			},
			limit:  1,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			markdown := ToMarkdown(test.input, test.limit)
			assert.Equal(t, test.output, markdown)
		})
	}
}
