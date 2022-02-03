package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConpassEventList_ToProfileMarkdown(t *testing.T) {
	publishedAt1, _ := time.Parse(time.RFC3339, "2022-02-01T14:59:00+00:00")
	publishedAt2, _ := time.Parse(time.RFC3339, "2022-02-01T15:00:00+00:00")
	organizer := "organizer"

	tests := map[string]struct {
		input  ConpassEventList
		output string
	}{
		"マークダウンに変換できる": {
			input: ConpassEventList{
				&connpassEvent{title: "イベントの例1", link: "https://example.com/1", ownerNickName: organizer, startedAt: publishedAt1},
				&connpassEvent{title: "イベントの例2", link: "https://example.com/2", ownerNickName: "unkwown", startedAt: publishedAt2},
			},
			output: "\n- Feb 1 **Organizer** [イベントの例1](https://example.com/1)\n- Feb 2 [イベントの例2](https://example.com/2)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			markdown := test.input.ToProfileMarkdown(organizer)
			assert.Equal(t, test.output, markdown)
		})
	}
}

func TestConpassEventList_SortByPublishedAt(t *testing.T) {
	first, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:01+09:00")
	second, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:02+09:00")
	third, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:03+09:00")

	tests := map[string]struct {
		input  ConpassEventList
		output ConpassEventList
	}{
		"開始時間の遅い順となる": {
			input: ConpassEventList{
				&connpassEvent{startedAt: second},
				&connpassEvent{startedAt: first},
				&connpassEvent{startedAt: third},
			},
			output: ConpassEventList{
				&connpassEvent{startedAt: third},
				&connpassEvent{startedAt: second},
				&connpassEvent{startedAt: first},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile := test.input.SortByPublishedAt()
			assert.Equal(t, test.output, profile)
		})
	}
}

func TestConpassEventList_Limit(t *testing.T) {
	tests := map[string]struct {
		input  ConpassEventList
		limit  int
		output ConpassEventList
	}{
		"イベント数を制限できる": {
			input: ConpassEventList{
				&connpassEvent{title: "first"},
				&connpassEvent{title: "second"},
				&connpassEvent{title: "third"},
			},
			limit: 1,
			output: ConpassEventList{
				&connpassEvent{title: "first"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile := test.input.Limit(test.limit)
			assert.Equal(t, test.output, profile)
		})
	}
}
