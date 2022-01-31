package usecase

import (
	"context"
	"fmt"
)

type UpdateProfileUsecase struct {
	profileIO  ProfileIO
	zennClient ZennClient
}

func (u UpdateProfileUsecase) Exec(ctx context.Context) error {
	readme, err := u.profileIO.Scan()
	if err != nil {
		return err
	}

	articles, err := u.zennClient.FetchArticles(ctx, "kumackey")
	if err != nil {
		return err
	}

	readme, err = readme.Replace(articles.SortByPublishedAt().ToProfileMarkdown())
	if err != nil {
		return err
	}

	err = u.profileIO.Write(readme)
	if err != nil {
		return err
	}

	fmt.Println(*readme)

	return nil
}

func NewUpdateProfileUsecase(profileIO ProfileIO, zennClient ZennClient) UpdateProfileUsecase {
	return UpdateProfileUsecase{
		profileIO:  profileIO,
		zennClient: zennClient,
	}
}
