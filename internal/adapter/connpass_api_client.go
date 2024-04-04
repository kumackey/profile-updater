package adapter

import (
	"context"

	"github.com/kumackey/profile-updater/internal/domain"
	"github.com/tenntenn/connpass"
)

type ConnpassAPIClient struct{}

func (c ConnpassAPIClient) FetchEventList(ctx context.Context, userNickname string) ([]domain.Item, error) {
	client := connpass.NewClient()

	// https://connpass.com/about/api/
	params, err := connpass.SearchParam(connpass.Nickname(userNickname))
	if err != nil {
		return nil, err
	}

	response, err := client.Search(ctx, params)
	if err != nil {
		return nil, err
	}

	list := make([]domain.Item, 0, len(response.Events))
	for _, event := range response.Events {
		list = append(list, domain.NewConpassEvent(event.Title, event.URL, event.StartedAt))
	}

	return list, nil
}
