package adapter

import (
	"bufio"
	"github.com/kumackey/qiita-profile/pkg/domain"
	"os"
)

const filenameReadMe = "README.md"

type ReadmeFileOS struct{}

func (s ReadmeFileOS) Scan() (*domain.Readme, error) {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []domain.Line
	for scanner.Scan() {
		lines = append(lines, domain.Line(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &domain.Readme{
		Content: lines,
	}, nil
}

func (s ReadmeFileOS) Write(readme *domain.Readme) error {
	f, err := os.Create(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range readme.Content {
		_, err := f.WriteString(string(line) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
