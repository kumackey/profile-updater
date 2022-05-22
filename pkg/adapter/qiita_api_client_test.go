package adapter

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/kumackey/profile-updater/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

func TestQiitaAPIClient_FetchArticleList(t *testing.T) {
	tests := map[string]struct {
		userID       string
		articleCount int
	}{
		"kumackeyは10記事以上書いている": {
			"kumackey", 10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qiita := QiitaAPIClient{}
			list, err := qiita.FetchArticleList(context.Background(), test.userID, 0)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(list), test.articleCount)
		})
	}
}

func TestQiitaAPIClient_FetchArticleList_Failed(t *testing.T) {
	//nolint:gosec // ランダム文字列を作りたいだけなので無視
	random := strconv.Itoa(rand.Intn(100000))

	tests := map[string]struct {
		userID string
	}{
		"適当なユーザ名ではフィードが発見できない": {
			"unknownUser" + random,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qiita := QiitaAPIClient{}
			_, err := qiita.FetchArticleList(context.Background(), test.userID, 0)
			assert.Equal(t, usecase.ErrQiitaAuthorNotFound, err)
		})
	}
}
