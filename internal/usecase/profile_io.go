package usecase

import (
	"github.com/kumackey/profile-updater/internal/domain"
)

type ProfileIO interface {
	Scan() (*domain.Profile, error)
	Write(*domain.Profile) error
}
