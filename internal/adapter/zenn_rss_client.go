package adapter

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/kumackey/profile-updater/internal/domain"
)

type ZennRSSClient struct{}

func (r ZennRSSClient) FetchArticleList(ctx context.Context, userID string) ([]domain.ZennArticle, error) {
	// https://zenn.dev/zenn/articles/zenn-feed-rss
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://zenn.dev/"+userID+"/feed", http.NoBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		if http.StatusInternalServerError < resp.StatusCode {
			return nil, domain.ErrZennInternalServerError
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, domain.ErrZennAuthorNotFound
		}

		return nil, domain.ErrZennUnknownError
	}

	var rss zennUserFeed
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&rss)
	if err != nil {
		return nil, err
	}

	// https://go-critic.com/overview#rangevalcopy
	list := make([]domain.ZennArticle, 0, len(rss.Items))
	for i := range rss.Items {
		publishedAt, err := time.Parse(time.RFC1123, rss.Items[i].PubDate)
		if err != nil {
			return nil, err
		}

		list = append(list, domain.NewZennArticle(rss.Items[i].Title, rss.Items[i].Link, publishedAt))
	}

	return list, nil
}
