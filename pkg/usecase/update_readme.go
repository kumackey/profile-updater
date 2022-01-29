package usecase

import (
	"bufio"
	"os"
)

const (
	filenameReadMe = "README.md"
	beginLine      = "<!-- begin line of qiita profile -->"
	endLine        = "<!-- end line of qiita profile -->"
)

type UpdateReadMe struct{}

func NewUpdateReadMe() *UpdateReadMe {
	return &UpdateReadMe{}
}

func (u *UpdateReadMe) Exec() error {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []line
	for scanner.Scan() {
		lines = append(lines, line(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	f, err = os.Create(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	lines = replaceLines(lines)
	for _, line := range lines {
		_, err := f.WriteString(string(line) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

type line string

func replaceLines(lines []line) []line {
	replacedLines := make([]line, 0, len(lines))
	writeMode := false
	for _, line := range lines {
		if line == endLine {
			replacedLines = append(replacedLines, "replaced line", line)
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

	return replacedLines
}
