package domain

type line string

type Profile struct {
	Content []line
}

const (
	beginLine = "<!-- begin line of profile updater -->"
	endLine   = "<!-- end line of profile updater -->"
)

func (p *Profile) Replace(values []string) *Profile {
	// TODO: 同じ行数しか確保できてない
	replacedLines := make([]line, 0, len(p.Content))
	writeMode := false
	for _, contentLine := range p.Content {
		if contentLine == endLine {
			lines := make([]line, 0, len(values))
			for _, value := range values {
				lines = append(lines, line(value))
			}

			replacedLines = append(replacedLines, lines...)
			replacedLines = append(replacedLines, contentLine)
			writeMode = false

			continue
		}

		if writeMode {
			continue
		}

		replacedLines = append(replacedLines, contentLine)

		if contentLine == beginLine {
			writeMode = true
		}
	}

	return &Profile{Content: replacedLines}
}

func NewProfile(values []string) *Profile {
	lines := make([]line, 0, len(values))
	for _, value := range values {
		lines = append(lines, line(value))
	}

	return &Profile{lines}
}
