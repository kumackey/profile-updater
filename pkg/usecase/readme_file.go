package usecase

import "github.com/kumackey/qiita-profile/pkg/domain"

type ProfileIO interface {
	Scan() (*domain.Profile, error)
	Write(*domain.Profile) error
}
