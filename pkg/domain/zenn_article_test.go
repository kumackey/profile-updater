package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZennArticle_ToMarkdown(t *testing.T) {
	tests := map[string]struct {
		input  *ZennArticle
		output string
	}{
		"マークダウンに変換できること": {
			input: &ZennArticle{
				Title: "タイトル",
				Link:  "https://example.com",
			},
			output: "[タイトル](https://example.com)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			markdown := test.input.ToMarkdown()
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
