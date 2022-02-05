package usecase

import (
	"context"
)

type UpdateProfileUsecase struct {
	profileIO      ProfileIO
	zennClient     ZennClient
	connpassClient ConnpassClient
}

// DefaultMaxLines はデフォルトでの最大行数
const DefaultMaxLines = 5

func (u UpdateProfileUsecase) Exec(ctx context.Context, zennUserID string, zennMaxArticles int, connpassNickname string, connpassMaxEvents int) error {
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

	if connpassNickname != "" {
		connpassList, err := u.connpassClient.FetchEventList(ctx, connpassNickname)
		if err != nil {
			return err
		}

		profile, err = profile.ReplaceConnpass(connpassList.SortByPublishedAt().Limit(connpassMaxEvents).ToProfileMarkdown(connpassNickname))
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
