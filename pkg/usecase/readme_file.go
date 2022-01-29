package usecase

import "github.com/kumackey/qiita-profile/pkg/domain"

type ReadmeFile interface {
	Scan() (*domain.Readme, error)
	Write(*domain.Readme) error
}
