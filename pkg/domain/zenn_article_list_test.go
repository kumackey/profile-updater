package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZennArticleList_ToProfileMarkdown(t *testing.T) {
	publishedAt1, _ := time.Parse(time.RFC3339, "2022-02-01T14:59:00+00:00")
	publishedAt2, _ := time.Parse(time.RFC3339, "2022-02-01T15:00:00+00:00")

	tests := map[string]struct {
		input  ZennArticleList
		output string
	}{
		"マークダウンに変換できる": {
			input: ZennArticleList{
				&zennArticle{
					title:       "記事の例1",
					link:        "https://example.com/1",
					publishedAt: publishedAt1,
				},
				&zennArticle{
					title:       "記事の例2",
					link:        "https://example.com/2",
					publishedAt: publishedAt2,
				},
			},

			output: "\n- Feb 1 [記事の例1](https://example.com/1)\n- Feb 2 [記事の例2](https://example.com/2)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			markdown := test.input.ToProfileMarkdown()
			assert.Equal(t, test.output, markdown)
		})
	}
}

func TestZennArticleList_SortByPublishedAt(t *testing.T) {
	first, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:01+09:00")
	second, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:02+09:00")
	third, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:03+09:00")

	tests := map[string]struct {
		input  ZennArticleList
		output ZennArticleList
	}{
		"出版の遅い順となる": {
			input: ZennArticleList{
				&zennArticle{publishedAt: second},
				&zennArticle{publishedAt: first},
				&zennArticle{publishedAt: third},
			},
			output: ZennArticleList{
				&zennArticle{publishedAt: third},
				&zennArticle{publishedAt: second},
				&zennArticle{publishedAt: first},
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

func TestZennArticleList_Limit(t *testing.T) {
	tests := map[string]struct {
		input  ZennArticleList
		limit  int
		output ZennArticleList
	}{
		"記事数を制限できる": {
			input: ZennArticleList{
				&zennArticle{title: "first"},
				&zennArticle{title: "second"},
				&zennArticle{title: "third"},
			},
			limit: 1,
			output: ZennArticleList{
				&zennArticle{title: "first"},
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
