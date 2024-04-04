package domain

import (
	"context"
)

type ConnpassClient interface {
	FetchEventList(ctx context.Context, userNickname string) (ConpassEventList, error)
}
