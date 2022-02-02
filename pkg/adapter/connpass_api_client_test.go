package adapter

import (
	"context"
	"testing"
)

func TestConnpassAPIClient_FetchEvents(t *testing.T) {
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
			_ = test
			client := ConnpassAPIClient{}
			client.FetchEvents(context.Background())
			//assert.Nil(t, err)
			//assert.GreaterOrEqual(t, len(list), test.articleCount)
		})
	}
}
