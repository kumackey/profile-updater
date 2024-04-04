package usecase

import "github.com/kumackey/profile-updater/pkg/domain"

type ProfileIO interface {
	Scan() (*domain.Profile, error)
	Write(*domain.Profile) error
}
