package adapter

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/kumackey/profile-updater/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestRSSClient_FetchItems(t *testing.T) {
	tests := map[string]struct {
		url       *url.URL
		wantCount int
	}{
		"kumackeyは8記事以上書いていること": {
			mustURL("https://zenn.dev/kumackey/feed"), 8,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := RSSClient{&http.Client{}}
			items, err := c.FetchItems(context.Background(), test.url)
			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(items), test.wantCount)
		})
	}
}

func mustURL(v string) *url.URL {
	u, err := url.Parse(v)
	if err != nil {
		panic(err)
	}
	return u
}

func TestRSSClient_FetchItems_Failed(t *testing.T) {
	tests := map[string]struct {
		url *url.URL
	}{
		"適当なユーザ名ではフィードが発見できないこと": {
			mustURL("https://zenn.dev/invalidinvalidusersampletest/feed"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := RSSClient{&http.Client{}}
			_, err := c.FetchItems(context.Background(), test.url)
			assert.Equal(t, domain.ErrRssNotFound, err)
		})
	}
}
