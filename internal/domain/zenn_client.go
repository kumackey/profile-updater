package domain

import (
	"context"
	"errors"
)

var (
	ErrZennAuthorNotFound      = errors.New("zenn author not found")
	ErrZennInternalServerError = errors.New("zenn internal server error")
	ErrZennUnknownError        = errors.New("zenn unknown error")
)

type ZennClient interface {
	FetchArticleList(ctx context.Context, userID string) ([]ZennArticle, error)
}
