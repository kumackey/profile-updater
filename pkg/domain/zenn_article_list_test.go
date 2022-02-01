package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZennArticleList_ToProfileMarkdown(t *testing.T) {
	tests := map[string]struct {
		input  ZennArticleList
		output string
	}{
		"マークダウンに変換できること": {
			input: ZennArticleList{
				&ZennArticle{
					Title: "記事の例1",
					Link:  "https://example.com/1",
				},
				&ZennArticle{
					Title: "記事の例2",
					Link:  "https://example.com/2",
				},
			},

			output: "\n- [記事の例1](https://example.com/1)\n- [記事の例2](https://example.com/2)\n",
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
		"出版の遅い順となること": {
			input: ZennArticleList{
				&ZennArticle{PublishedAt: second},
				&ZennArticle{PublishedAt: first},
				&ZennArticle{PublishedAt: third},
			},
			output: ZennArticleList{
				&ZennArticle{PublishedAt: third},
				&ZennArticle{PublishedAt: second},
				&ZennArticle{PublishedAt: first},
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
		"記事数を制限できること": {
			input: ZennArticleList{
				&ZennArticle{Title: "first"},
				&ZennArticle{Title: "second"},
				&ZennArticle{Title: "third"},
			},
			limit: 1,
			output: ZennArticleList{
				&ZennArticle{Title: "first"},
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
