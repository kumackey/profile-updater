package usecase

import (
	"context"
)

type UpdateProfileUsecase struct {
	profileIO      ProfileIO
	zennClient     ZennClient
	connpassClient ConnpassClient
}

func (u UpdateProfileUsecase) Exec(ctx context.Context, zennUserID string, zennMaxArticles int, connpassNickName string, connpassMaxEvents int) error {
	profile, err := u.profileIO.Scan()
	if err != nil {
		return err
	}

	if zennUserID != "" {
		zennList, err := u.zennClient.FetchArticleList(ctx, zennUserID)
		if err != nil {
			return err
		}

		profile, err = profile.ReplaceZenn(zennList.SortByPublishedAt().Limit(zennMaxArticles).ToProfileMarkdown())
		if err != nil {
			return err
		}
	}

	if connpassNickName != "" {
		connpassList, err := u.connpassClient.FetchEventList(ctx, connpassNickName)
		if err != nil {
			return err
		}

		profile, err = profile.ReplaceConnpass(connpassList.SortByPublishedAt().Limit(connpassMaxEvents).ToProfileMarkdown(connpassNickName))
		if err != nil {
			return err
		}
	}

	err = u.profileIO.Write(profile)
	if err != nil {
		return err
	}

	return nil
}

func NewUpdateProfileUsecase(profileIO ProfileIO, zennClient ZennClient, connpassClient ConnpassClient) UpdateProfileUsecase {
	return UpdateProfileUsecase{
		profileIO:      profileIO,
		zennClient:     zennClient,
		connpassClient: connpassClient,
	}
}
