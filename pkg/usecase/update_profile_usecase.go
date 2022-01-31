package usecase

import (
	"context"
	"fmt"
)

type UpdateProfileUsecase struct {
	ProfileIO  ProfileIO
	ZennClient ZennClient
}

func (u UpdateProfileUsecase) Exec() error {
	readme, err := u.ProfileIO.Scan()
	if err != nil {
		return err
	}

	ctx := context.Background()
	articles, err := u.ZennClient.FetchArticles(ctx, "kumackey")
	if err != nil {
		return err
	}

	readme, err = readme.Replace(articles.SortByPublishedAt().ToProfileMarkdown())
	if err != nil {
		return err
	}

	err = u.ProfileIO.Write(readme)
	if err != nil {
		return err
	}

	fmt.Println(*readme)

	return nil
}
