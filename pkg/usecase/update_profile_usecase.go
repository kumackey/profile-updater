package usecase

import (
	"context"
)

// DefaultMaxLines はデフォルトでの最大行数
const DefaultMaxLines = 5

type UpdateProfileUsecase struct {
	profileIO      ProfileIO
	zennClient     ZennClient
	connpassClient ConnpassClient
	qiitaClient    QiitaClient
}

type UpdateProfileUsecaseInput struct {
	zennUserID        string
	zennMaxArticles   int
	connpassNickname  string
	connpassMaxEvents int
	qiitaUserID       string
	qiitaMaxArticles  int
}

func (u UpdateProfileUsecase) Exec(ctx context.Context, input UpdateProfileUsecaseInput) error {
	profile, err := u.profileIO.Scan()
	if err != nil {
		return err
	}

	if input.zennUserID != "" {
		zennList, err := u.zennClient.FetchArticleList(ctx, input.zennUserID)
		if err != nil {
			return err
		}

		replaceValue := zennList.SortByPublishedAt().Limit(input.zennMaxArticles).ToProfileMarkdown()

		profile, err = profile.ReplaceZenn(replaceValue)
		if err != nil {
			return err
		}
	}

	if input.connpassNickname != "" {
		connpassList, err := u.connpassClient.FetchEventList(ctx, input.connpassNickname)
		if err != nil {
			return err
		}

		replaceValue := connpassList.SortByPublishedAt().
			Limit(input.connpassMaxEvents).
			ToProfileMarkdown(input.connpassNickname)

		profile, err = profile.ReplaceConnpass(replaceValue)
		if err != nil {
			return err
		}
	}

	if input.qiitaUserID != "" {
		connpassList, err := u.qiitaClient.FetchArticleList(ctx, input.qiitaUserID)
		if err != nil {
			return err
		}

		replaceValue := connpassList.SortByPublishedAt().Limit(input.qiitaMaxArticles).ToProfileMarkdown()

		profile, err = profile.ReplaceConnpass(replaceValue)
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

func NewUpdateProfileUsecase(
	profileIO ProfileIO, zennClient ZennClient, connpassClient ConnpassClient, qiitaClient QiitaClient,
) UpdateProfileUsecase {
	return UpdateProfileUsecase{
		profileIO:      profileIO,
		zennClient:     zennClient,
		connpassClient: connpassClient,
		qiitaClient:    qiitaClient,
	}
}

func NewUpdateProfileUseCaseInput(
	zennUserID string,
	zennMaxArticles int,
	connpassNickname string,
	connpassMaxEvents int,
	qiitaUserID string,
	qiitaMaxArticles int,
) UpdateProfileUsecaseInput {
	return UpdateProfileUsecaseInput{
		zennUserID:        zennUserID,
		zennMaxArticles:   zennMaxArticles,
		connpassNickname:  connpassNickname,
		connpassMaxEvents: connpassMaxEvents,
		qiitaUserID:       qiitaUserID,
		qiitaMaxArticles:  qiitaMaxArticles,
	}
}
