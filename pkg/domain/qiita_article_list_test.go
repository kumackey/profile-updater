package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQiitaArticleList_ToProfileMarkdown(t *testing.T) {
	publishedAt1, _ := time.Parse(time.RFC3339, "2022-02-01T14:59:00+00:00")
	publishedAt2, _ := time.Parse(time.RFC3339, "2022-02-01T15:00:00+00:00")

	tests := map[string]struct {
		input  QiitaArticleList
		output string
	}{
		"マークダウンに変換できる": {
			input: QiitaArticleList{
				&qiitaArticle{
					title:       "記事の例1",
					link:        "https://example.com/1",
					publishedAt: publishedAt1,
				},
				&qiitaArticle{
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

func TestQiitaArticleList_SortByPublishedAt(t *testing.T) {
	first, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:01+09:00")
	second, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:02+09:00")
	third, _ := time.Parse(time.RFC3339, "2022-02-01T00:00:03+09:00")

	tests := map[string]struct {
		input  QiitaArticleList
		output QiitaArticleList
	}{
		"出版の遅い順となる": {
			input: QiitaArticleList{
				&qiitaArticle{publishedAt: second},
				&qiitaArticle{publishedAt: first},
				&qiitaArticle{publishedAt: third},
			},
			output: QiitaArticleList{
				&qiitaArticle{publishedAt: third},
				&qiitaArticle{publishedAt: second},
				&qiitaArticle{publishedAt: first},
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
