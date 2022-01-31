package adapter

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/kumackey/profile-updater/pkg/domain"
)

type ZennRSS struct{}

func (c ZennRSS) FetchArticles(ctx context.Context, userID string) (domain.ZennArticles, error) {
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

	var rss rssXML
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&rss)
	if err != nil {
		return nil, err
	}

	// https://go-critic.com/overview#rangevalcopy
	var articles domain.ZennArticles
	for i := range rss.Items {
		publishedAt, err := time.Parse(time.RFC1123, rss.Items[i].PubDate)
		if err != nil {
			return nil, err
		}
		article := domain.ZennArticle{
			Title:       rss.Items[i].Title,
			Link:        rss.Items[i].Link,
			EnClosure:   domain.EnClosure{URL: rss.Items[i].Enclosure.URL},
			PublishedAt: publishedAt,
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

// https://zenn.dev/link/comments/731a3eba374a8e
type rssXML struct {
	XMLName xml.Name `xml:"rss"`
	Items   []struct {
		Title       string `xml:"title"`
		Description struct {
			XMLName xml.Name `xml:"description"`
			Value   string   `xml:",cdata"`
		}
		Link string `xml:"link"`
		GuID struct {
			XMLName     xml.Name `xml:"guid"`
			IsPermaLink bool     `xml:"isPermaLink,attr"`
			Value       string   `xml:",chardata"`
		}
		PubDate   string `xml:"pubDate"`
		Enclosure struct {
			XMLName string `xml:"enclosure"`
			URL     string `xml:"url,attr"`
			Length  int    `xml:"length,attr"`
			Type    string `xml:"type,attr"`
		}
		Creator string `xml:"creator"`
	} `xml:"channel>item"`
}
