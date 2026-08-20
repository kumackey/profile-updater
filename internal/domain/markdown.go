package domain

import (
	"net/url"
	"strings"
)

// 外部サービスから取得したタイトルは信頼できないため、Markdownとして無害化する。
// リンク記法の偽装(`](https://evil.example)`)や、HTMLタグ・置き換えマーカーのコメントの混入を防ぐ。
var markdownTextEscaper = strings.NewReplacer(
	`\`, `\\`,
	"[", `\[`,
	"]", `\]`,
	"<", `\<`,
	">", `\>`,
	"\r\n", " ",
	"\r", " ",
	"\n", " ",
)

func escapeMarkdownText(text string) string {
	return markdownTextEscaper.Replace(text)
}

var markdownLinkEscaper = strings.NewReplacer(
	"(", "%28",
	")", "%29",
	" ", "%20",
)

// sanitizeLink はhttp/https以外のリンクを空文字にし、Markdownのリンク記法を壊す文字をエスケープする。
func sanitizeLink(link string) string {
	u, err := url.Parse(strings.TrimSpace(link))
	if err != nil {
		return ""
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return ""
	}

	return markdownLinkEscaper.Replace(u.String())
}

// toMarkdownLink は外部由来のタイトルとリンクから、安全なMarkdownのリンクを組み立てる。
func toMarkdownLink(title, link string) string {
	return "[" + escapeMarkdownText(title) + "](" + sanitizeLink(link) + ")"
}
