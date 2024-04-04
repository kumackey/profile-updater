package usecase

import (
	"context"

	"github.com/kumackey/profile-updater/internal/domain"
)

type ConnpassClient interface {
	FetchEventList(ctx context.Context, userNickname string) (domain.ConpassEventList, error)
}
