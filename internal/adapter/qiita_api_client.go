package adapter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kumackey/profile-updater/internal/domain"
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

func (c QiitaAPIClient) FetchArticleList(
	ctx context.Context,
	userID string,
	limit int,
) ([]domain.QiitaArticle, error) {
	client := &http.Client{}

	// https://qiita.com/api/v2/docs#%E6%8A%95%E7%A8%BF
	url := "https://qiita.com/api/v2/items?per_page=" + strconv.Itoa(limit) + "&query=user:" + userID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		if http.StatusInternalServerError < resp.StatusCode {
			return nil, domain.ErrQiitaInternalServerError
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, domain.ErrQiitaAuthorNotFound
		}

		return nil, domain.ErrQiitaUnknownError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []qiitaAPIItem
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		// 投稿がゼロの可能性も考慮する必要があるが、そんなユーザはこのActionsを使わない理論により無視する
		return nil, domain.ErrQiitaAuthorNotFound
	}

	list := make([]domain.QiitaArticle, 0, len(data))
	for i := range data {
		list = append(list, domain.NewQiitaArticle(data[i].Title, data[i].URL, data[i].LikeCount, data[i].CreatedAt))
	}

	return list, nil
}
