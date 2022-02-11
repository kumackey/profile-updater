package usecase

import (
	"context"
)

type UpdateProfileUsecase struct {
	profileIO      ProfileIO
	zennClient     ZennClient
	connpassClient ConnpassClient
	qiitaClient    QiitaClient
}

// DefaultMaxLines はデフォルトでの最大行数
const DefaultMaxLines = 5

func (u UpdateProfileUsecase) Exec(
	ctx context.Context, zennUserID string, zennMaxArticles int,
	connpassNickname string, connpassMaxEvents int,
	qiitaUserID string, qiitaMaxArticles int,
) error {
	profile, err := u.profileIO.Scan()
	if err != nil {
		return err
	}

	if zennUserID != "" {
		zennList, err := u.zennClient.FetchArticleList(ctx, zennUserID)
		if err != nil {
			return err
		}

		replaceValue := zennList.SortByPublishedAt().Limit(zennMaxArticles).ToProfileMarkdown()

		profile, err = profile.ReplaceZenn(replaceValue)
		if err != nil {
			return err
		}
	}

	if connpassNickname != "" {
		connpassList, err := u.connpassClient.FetchEventList(ctx, connpassNickname)
		if err != nil {
			return err
		}

		replaceValue := connpassList.SortByPublishedAt().Limit(connpassMaxEvents).ToProfileMarkdown(connpassNickname)

		profile, err = profile.ReplaceConnpass(replaceValue)
		if err != nil {
			return err
		}
	}

	if qiitaUserID != "" {
		connpassList, err := u.qiitaClient.FetchArticleList(ctx, qiitaUserID)
		if err != nil {
			return err
		}

		replaceValue := connpassList.SortByPublishedAt().Limit(qiitaMaxArticles).ToProfileMarkdown()

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
