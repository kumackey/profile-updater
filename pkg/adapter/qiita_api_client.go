package adapter

import (
	"context"
	"encoding/json"
	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/kumackey/profile-updater/pkg/usecase"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type QiitaAPIClient struct{}

type qiitaAPIItem struct {
	LikeCount int    `json:"likes_count"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Tags      []struct {
		Name string `json:"name"`
	} `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
}

func (c QiitaAPIClient) FetchArticleList(ctx context.Context, userID string, limit int) (domain.QiitaArticleList, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://qiita.com/api/v2/items?per_page="+strconv.Itoa(limit)+"&query=qiita+user:"+userID, http.NoBody)
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
			return nil, usecase.ErrQiitaInternalServerError
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, usecase.ErrQiitaAuthorNotFound
		}

		return nil, usecase.ErrQiitaUnknownError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []qiitaAPIItem
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	list := make(domain.QiitaArticleList, 0, len(data))
	for _ = range data {
		//publishedAt, err := time.Parse(time.RFC3339, atom.Entries[i].Published)
		//if err != nil {
		//	return nil, err
		//}
		//
		//list = append(list, domain.NewQiitaArticle(atom.Entries[i].Title, atom.Entries[i].URL, publishedAt))
	}

	return list, nil
}
