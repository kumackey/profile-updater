package domain

import (
	"context"
	"errors"
)

var (
	ErrQiitaAuthorNotFound      = errors.New("qiita author not found")
	ErrQiitaInternalServerError = errors.New("qiita internal server error")
	ErrQiitaUnknownError        = errors.New("qiita unknown error")
)

type QiitaClient interface {
	FetchArticleList(ctx context.Context, userID string, limit int) ([]QiitaArticle, error)
}
