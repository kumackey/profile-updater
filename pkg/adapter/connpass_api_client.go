package adapter

import (
	"context"
	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/tenntenn/connpass"
)

type ConnpassAPIClient struct{}

func (c ConnpassAPIClient) FetchEventList(ctx context.Context, userNickName string) (domain.ConpassEventList, error) {
	client := connpass.NewClient()

	// https://connpass.com/about/api/
	params, err := connpass.SearchParam(connpass.Nickname(userNickName))
	if err != nil {
		return nil, err
	}

	response, err := client.Search(ctx, params)
	if err != nil {
		return nil, err
	}

	list := make(domain.ConpassEventList, 0, len(response.Events))
	for _, event := range response.Events {
		list = append(list, domain.NewConpassEvent(event.Title, event.URL, event.OwnerNickname, event.StartedAt))
	}

	return list, nil
}
