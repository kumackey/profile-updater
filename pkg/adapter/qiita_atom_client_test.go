package adapter

import (
	"context"
	"github.com/kumackey/profile-updater/pkg/usecase"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

func TestQiitaAtomClient_FetchArticleList(t *testing.T) {
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
			qiita := QiitaAtomClient{}
			list, err := qiita.FetchArticleList(context.Background(), test.userID)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(list), test.articleCount)
		})
	}
}

func TestQiitaAtomClient_FetchArticleList_Failed(t *testing.T) {
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
			qiita := QiitaAtomClient{}
			_, err := qiita.FetchArticleList(context.Background(), test.userID)
			assert.Equal(t, usecase.ErrQiitaAuthorNotFound, err)
		})
	}
}
