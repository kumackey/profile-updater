package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZennArticle_ToProfileMarkdown(t *testing.T) {
	tests := map[string]struct {
		input  ZennArticles
		output string
	}{
		"マークダウンに変換できること": {
			input: ZennArticles{
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

func TestZennArticles_SortByPublishedAt(t *testing.T) {
	first, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:01+09:00")
	second, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:02+09:00")
	third, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:03+09:00")

	tests := map[string]struct {
		input  ZennArticles
		output ZennArticles
	}{
		"出版の遅い順となること": {
			input: ZennArticles{
				&ZennArticle{PublishedAt: second},
				&ZennArticle{PublishedAt: first},
				&ZennArticle{PublishedAt: third},
			},
			output: ZennArticles{
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
