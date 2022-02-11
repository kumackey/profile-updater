package usecase

import (
	"context"
	"errors"

	"github.com/kumackey/profile-updater/pkg/domain"
)

var (
	ErrQiitaAuthorNotFound      = errors.New("qiita author not found")
	ErrQiitaInternalServerError = errors.New("qiita internal server error")
	ErrQiitaUnknownError        = errors.New("qiita unknown error")
)

type QiitaClient interface {
	FetchArticleList(ctx context.Context, userID string) (domain.QiitaArticleList, error)
}
