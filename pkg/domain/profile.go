package domain

import (
	"errors"
)

type Profile struct {
	Content []line
}

const (
	beginLine = "<!-- begin line of profile updater -->"
	endLine   = "<!-- end line of profile updater -->"
)

var ErrReplaceLinesNotFound = errors.New("replace lines not found")

func (p *Profile) Replace(values []string) (*Profile, error) {
	// TODO: 同じ行数しか確保できてない
	replacedLines := make([]line, 0, len(p.Content))

	isReplaced := false

	for _, contentLine := range p.Content {
		if contentLine.shouldReplace {
			lines := make([]line, 0, len(values))
			for _, value := range values {
				lines = append(lines, line{value, false})
			}

			replacedLines = append(replacedLines, line{beginLine, false})
			replacedLines = append(replacedLines, lines...)
			replacedLines = append(replacedLines, line{endLine, false})
			isReplaced = true

			continue
		}

		replacedLines = append(replacedLines, contentLine)
	}
	profile := &Profile{Content: replacedLines}
	if !isReplaced {
		return profile, ErrReplaceLinesNotFound
	}

	return profile, nil
}

func NewProfile(values []string) *Profile {
	lines := make([]line, 0, len(values))
	isReplaced := false
	for _, value := range values {
		if value == beginLine {
			isReplaced = true
			lines = append(lines, line{"", true})
			continue
		}

		if isReplaced {
			continue
		}

		lines = append(lines, line{value, isReplaced})

		if value == endLine {
			isReplaced = false
			continue
		}
	}

	return &Profile{lines}
}
