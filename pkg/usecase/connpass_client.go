package usecase

import (
	"context"
	"github.com/kumackey/profile-updater/pkg/domain"
)

type ConnpassClient interface {
	FetchEventList(ctx context.Context, userNickName string) (domain.ConpassEventList, error)
}
