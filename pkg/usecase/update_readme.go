package usecase

import (
	"bufio"
	"os"

	"github.com/kumackey/qiita-profile/pkg/domain"
)

const (
	filenameReadMe = "README.md"
	beginLine      = "<!-- begin line of qiita profile -->"
	endLine        = "<!-- end line of qiita profile -->"
)

type UpdateReadme struct {
}

func NewUpdateReadMe() *UpdateReadme {
	return &UpdateReadme{}
}

func (u *UpdateReadme) Exec() error {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []domain.Line
	for scanner.Scan() {
		lines = append(lines, domain.Line(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	f, err = os.Create(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	readme := &domain.Readme{
		Content: lines,
	}

	readme = replaceLines(readme)
	for _, line := range readme.Content {
		_, err := f.WriteString(string(line) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func replaceLines(readme *domain.Readme) *domain.Readme {
	replacedLines := make([]domain.Line, 0, len(readme.Content))
	writeMode := false
	for _, line := range readme.Content {
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

	return &domain.Readme{Content: replacedLines}
}
