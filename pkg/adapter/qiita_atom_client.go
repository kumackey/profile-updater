package adapter

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

type QiitaAtomClient struct{}

func (r QiitaAtomClient) FetchArticleList(ctx context.Context, userID string) (domain.QiitaArticleList, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://qiita.com/"+userID+"/feed", http.NoBody)
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

	var atom qiitaUserFeed
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&atom)
	if err != nil {
		return nil, err
	}

	// https://go-critic.com/overview#rangevalcopy
	list := make(domain.QiitaArticleList, 0, len(atom.Entries))
	for i := range atom.Entries {
		publishedAt, err := time.Parse(time.RFC3339, atom.Entries[i].Published)
		if err != nil {
			return nil, err
		}

		list = append(list, domain.NewQiitaArticle(atom.Entries[i].Title, atom.Entries[i].URL, publishedAt))
	}

	return list, nil
}
