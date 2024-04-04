package adapter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnpassAPIClient_FetchEventList(t *testing.T) {
	tests := map[string]struct {
		userNickname string
		eventCount   int
	}{
		"kumackeyは10イベント以上は参加している": {
			"kumackey", 10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := ConnpassAPIClient{}
			list, err := client.FetchEventList(context.Background(), test.userNickname)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(list), test.eventCount)
		})
	}
}
