package adapter

import (
	"context"
	"github.com/kumackey/profile-updater/pkg/usecase"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZennRSS_FetchArticles(t *testing.T) {
	tests := map[string]struct {
		userID       string
		articleCount int
	}{
		"kumackeyは8記事以上書いている": {
			"kumackey", 8,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			zenn := ZennRSS{}
			articles, err := zenn.FetchArticles(context.Background(), test.userID)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(articles), test.articleCount)
		})
	}
}

func TestZennRSS_FetchArticles_Failed(t *testing.T) {
	random := strconv.Itoa(rand.Intn(100000))

	tests := map[string]struct {
		userID       string
		articleCount int
	}{
		"適当なユーザ名ではフィードが発見できない": {
			"unknownUser" + random, 8,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			zenn := ZennRSS{}
			_, err := zenn.FetchArticles(context.Background(), test.userID)
			assert.Equal(t, usecase.ErrZennAuthorNotFound, err)
		})
	}
}
