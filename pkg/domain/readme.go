package domain

type Line string

type Readme struct {
	Content []Line
}

const (
	beginLine = "<!-- begin line of qiita profile -->"
	endLine   = "<!-- end line of qiita profile -->"
)

func (r *Readme) Replace(lines []Line) *Readme {
	replacedLines := make([]Line, 0, len(r.Content))
	writeMode := false
	for _, line := range r.Content {
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

	return &Readme{Content: replacedLines}
}
