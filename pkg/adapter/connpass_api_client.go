package adapter

import (
	"context"
	"fmt"
	"github.com/tenntenn/connpass"
	"log"
)

type ConnpassAPIClient struct{}

func (c ConnpassAPIClient) FetchEvents(ctx context.Context) error {
	client := connpass.NewClient()
	params, err := connpass.SearchParam(connpass.Nickname("kumackey"))
	if err != nil {
		log.Fatal(err)
	}
	r, err := client.Search(ctx, params)

	for _, e := range r.Events {
		fmt.Println(e.Title)
	}

	return err
}
