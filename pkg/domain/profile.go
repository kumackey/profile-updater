package domain

type Line string

type Profile struct {
	Content []Line
}

const (
	beginLine = "<!-- begin line of profile updater -->"
	endLine   = "<!-- end line of profile updater -->"
)

func (p *Profile) Replace(lines []Line) *Profile {
	// TODO: 同じ行数しか確保できてない
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
