package adapter

import (
	"os"

	"github.com/kumackey/profile-updater/pkg/domain"
)

const (
	filenameReadMe  = "README.md"
	writePermission = 0o666
)

type ReadmeFileOS struct{}

func (s ReadmeFileOS) Scan() (*domain.Profile, error) {
	b, err := os.ReadFile(filenameReadMe)
	if err != nil {
		return nil, err
	}

	return domain.NewProfile(string(b)), nil
}

func (s ReadmeFileOS) Write(readme *domain.Profile) error {
	err := os.WriteFile(filenameReadMe, []byte(readme.Content), writePermission)
	if err != nil {
		return err
	}

	return nil
}
