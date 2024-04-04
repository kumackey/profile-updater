package adapter

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/kumackey/profile-updater/internal/domain"
)

type RSSClient struct {
	client *http.Client
}

func NewRssClient(client *http.Client) RSSClient {
	return RSSClient{client: client}
}

func (r RSSClient) FetchItems(ctx context.Context, url *url.URL) ([]domain.RssItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		if http.StatusInternalServerError < resp.StatusCode {
			return nil, domain.ErrRssInternalServerError
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, domain.ErrRssNotFound
		}

		return nil, domain.ErrRssUnknownError
	}

	var rss rssFeed
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&rss)
	if err != nil {
		return nil, fmt.Errorf("failed to decode rss: %w", err)
	}

	items := make([]domain.RssItem, 0, len(rss.Items))
	for _, item := range rss.Items {
		publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time: %w", err)
		}

		items = append(items, domain.NewRssItem(item.Title, item.Link, publishedAt))
	}

	return items, nil
}

type rssFeed struct {
	XMLName xml.Name  `xml:"rss"`
	Items   []rssItem `xml:"channel>item"`
}

type rssItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}
