package adapter

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/kumackey/profile-updater/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

func TestQiitaAtomClient_FetchArticleList(t *testing.T) {
	tests := map[string]struct {
		userID string
		limit  int
	}{
		"kumackeyは3記事は書いている": {
			"kumackey", 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qiita := QiitaAtomClient{}
			list, err := qiita.FetchArticleList(context.Background(), test.userID, test.limit)
			assert.Nil(t, err)
			assert.Equal(t, len(list), test.limit)
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
			_, err := qiita.FetchArticleList(context.Background(), test.userID, 0)
			assert.Equal(t, usecase.ErrQiitaAuthorNotFound, err)
		})
	}
}
