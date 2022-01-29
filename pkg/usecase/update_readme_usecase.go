package usecase

import (
	"github.com/kumackey/qiita-profile/pkg/domain"
)

const (
	beginLine = "<!-- begin line of qiita profile -->"
	endLine   = "<!-- end line of qiita profile -->"
)

type UpdateReadmeUsecase struct {
	ReadmeFile ReadmeFile
}

func (u UpdateReadmeUsecase) Exec() error {
	readme, err := u.ReadmeFile.Scan()
	if err != nil {
		return err
	}

	readme = replaceLines(readme)
	err = u.ReadmeFile.Write(readme)
	if err != nil {
		return err
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
