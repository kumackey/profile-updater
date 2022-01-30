package adapter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_Replace_Failed(t *testing.T) {
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
			zenn := ZennFeed{}
			articles, err := zenn.FetchArticles(context.Background(), test.userID)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(articles), test.articleCount)
		})
	}
}
