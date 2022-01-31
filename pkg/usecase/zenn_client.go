package usecase

import (
	"context"

	"github.com/kumackey/profile-updater/pkg/domain"
)

type ZennClient interface {
	FetchArticles(ctx context.Context, userID string) (domain.ZennArticles, error)
}
