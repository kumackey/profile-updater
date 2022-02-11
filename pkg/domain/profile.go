package domain

import (
	"errors"
	"regexp"
	"strings"
)

type Profile struct {
	Content string
}

const (
	newLineCode        = "q7ldEirrDId1o2crkGhA"
	regexZennBegin     = "<!-- profile updater begin: zenn -->"
	regexZennEnd       = "<!-- profile updater end: zenn -->"
	regexConnpassBegin = "<!-- profile updater begin: connpass -->" //nolint:gosec // セキュリティの問題でないので無視
	regexConnpassEnd   = "<!-- profile updater end: connpass -->"   //nolint:gosec // セキュリティの問題でないので無視
	regexQiitaBegin    = "<!-- profile updater begin: qiita -->"
	regexQiitaEnd      = "<!-- profile updater end: qiita -->"
)

var (
	ErrReplaceStatementNotFound = errors.New("replace statement not found")
)

func (p *Profile) ReplaceZenn(value string) (*Profile, error) {
	return p.replace(value, regexZennBegin, regexZennEnd)
}

func (p *Profile) ReplaceConnpass(value string) (*Profile, error) {
	return p.replace(value, regexConnpassBegin, regexConnpassEnd)
}

func (p *Profile) ReplaceQiita(value string) (*Profile, error) {
	return p.replace(value, regexQiitaBegin, regexQiitaEnd)
}

func (p *Profile) replace(value, replaceBegin, replaceEnd string) (*Profile, error) {
	re := regexp.MustCompile(replaceBegin + "(.*)" + replaceEnd)

	// 正規表現における改行コード対策
	// https://qiita.com/spiegel-im-spiegel/items/f1cc014ecb233afaa8af
	newLineReplaced := strings.NewReplacer(
		"\r\n", newLineCode,
		"\r", newLineCode,
		"\n", newLineCode,
	).Replace(p.Content)
	if !re.MatchString(newLineReplaced) {
		return nil, ErrReplaceStatementNotFound
	}

	replaced := re.ReplaceAllString(newLineReplaced, replaceBegin+value+replaceEnd)
	p.Content = strings.ReplaceAll(replaced, newLineCode, "\n")

	return p, nil
}

func NewProfile(content string) *Profile {
	return &Profile{content}
}
