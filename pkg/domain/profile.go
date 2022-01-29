package domain

type Line string

type Profile struct {
	Content []Line
}

const (
	beginLine = "<!-- begin line of qiita profile -->"
	endLine   = "<!-- end line of qiita profile -->"
)

func (p *Profile) Replace(lines []Line) *Profile {
	replacedLines := make([]Line, 0, len(p.Content))
	writeMode := false
	for _, line := range p.Content {
		if line == endLine {
			replacedLines = append(replacedLines, lines...)
			replacedLines = append(replacedLines, line)
			writeMode = false

			continue
		}

		if writeMode {
			continue
		}

		replacedLines = append(replacedLines, line)

		if line == beginLine {
			writeMode = true
		}
	}

	return &Profile{Content: replacedLines}
}
