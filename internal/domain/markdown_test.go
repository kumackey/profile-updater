package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMarkdownLink(t *testing.T) {
	tests := map[string]struct {
		title  string
		link   string
		output string
	}{
		"通常のタイトルとリンクはそのまま出力される": {
			"記事の例", "https://example.com/1", "[記事の例](https://example.com/1)",
		},
		"タイトルによるリンク偽装を防ぐ": {
			"foo](https://evil.example) [bar", "https://example.com/1",
			`[foo\](https://evil.example) \[bar](https://example.com/1)`,
		},
		"タイトル内のHTMLタグをエスケープする": {
			"<img src=x onerror=alert(1)>", "https://example.com/1",
			`[\<img src=x onerror=alert(1)\>](https://example.com/1)`,
		},
		"タイトル内の置き換えマーカーを無害化する": {
			regexZennEnd, "https://example.com/1",
			`[\<!-- profile updater end: zenn --\>](https://example.com/1)`,
		},
		"タイトル内の改行を空白に変換する": {
			"1行目\n2行目", "https://example.com/1", "[1行目 2行目](https://example.com/1)",
		},
		"http/https以外のリンクを除去する": {
			"記事の例", "javascript:alert(1)", "[記事の例]()",
		},
		"リンク内の括弧をエスケープする": {
			"記事の例", "https://example.com/a(b)c", "[記事の例](https://example.com/a%28b%29c)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.output, toMarkdownLink(test.title, test.link))
		})
	}
}
