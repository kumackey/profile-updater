package usecase

import (
	"context"
)

type UpdateProfileUsecase struct {
	profileIO  ProfileIO
	zennClient ZennClient
}

func (u UpdateProfileUsecase) Exec(ctx context.Context, zennUserID string, zennMaxArticles int) error {
	readme, err := u.profileIO.Scan()
	if err != nil {
		return err
	}

	articles, err := u.zennClient.FetchArticles(ctx, zennUserID)
	if err != nil {
		return err
	}

	readme, err = readme.Replace(articles.SortByPublishedAt().Limit(zennMaxArticles).ToProfileMarkdown())
	if err != nil {
		return err
	}

	err = u.profileIO.Write(readme)
	if err != nil {
		return err
	}

	return nil
}

func NewUpdateProfileUsecase(profileIO ProfileIO, zennClient ZennClient) UpdateProfileUsecase {
	return UpdateProfileUsecase{
		profileIO:  profileIO,
		zennClient: zennClient,
	}
}
