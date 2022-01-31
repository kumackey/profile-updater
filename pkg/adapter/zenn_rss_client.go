package adapter

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

type ZennRSSClient struct{}

func (r ZennRSSClient) FetchArticles(ctx context.Context, userID string) (domain.ZennArticles, error) {
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
			return nil, usecase.ErrZennInternalServerError
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, usecase.ErrZennAuthorNotFound
		}

		return nil, usecase.ErrZennUnknownError
	}

	var rss zennRSS
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&rss)
	if err != nil {
		return nil, err
	}

	// https://go-critic.com/overview#rangevalcopy
	articles := make(domain.ZennArticles, 0, len(rss.Items))
	for i := range rss.Items {
		article, err := r.convertItemToArticle(&rss.Items[i])
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (r ZennRSSClient) convertItemToArticle(item *zennItem) (*domain.ZennArticle, error) {
	publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
	if err != nil {
		return nil, err
	}

	return &domain.ZennArticle{
		Title:       item.Title,
		Link:        item.Link,
		EnClosure:   domain.EnClosure{URL: item.Enclosure.URL},
		PublishedAt: publishedAt,
	}, nil
}
