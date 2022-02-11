package adapter

import (
	"context"
	"github.com/stretchr/testify/assert"
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
