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
	newLineCode = "q7ldEirrDId1o2crkGhA"
	regexBegin  = "<!-- profile updater begin: zenn -->"
	regexEnd    = "<!-- profile updater end: zenn -->"
)

var (
	re = regexp.MustCompile(regexBegin + "(.*)" + regexEnd)

	ErrReplaceStatementNotFound = errors.New("replace statement not found")
)

func (p *Profile) Replace(value string) (*Profile, error) {
	// 正規表現における改行コード対策
	newLineReplaced := strings.NewReplacer(
		"\r\n", newLineCode,
		"\r", newLineCode,
		"\n", newLineCode,
	).Replace(p.Content)
	if !re.MatchString(newLineReplaced) {
		return nil, ErrReplaceStatementNotFound
	}

	replaced := re.ReplaceAllString(newLineReplaced, regexBegin+value+regexEnd)
	p.Content = strings.ReplaceAll(replaced, newLineCode, "\n")

	return p, nil
}

func NewProfile(content string) *Profile {
	return &Profile{content}
}
