package domain

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// DefaultMaxLines はデフォルトでの最大行数
const DefaultMaxLines = 5

type UpdateProfileUsecase struct {
	profileIO      ProfileIO
	connpassClient ConnpassClient
	qiitaClient    QiitaClient
	rssClient      RssClient
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
		// https://zenn.dev/zenn/articles/zenn-feed-rss
		zu, err := url.Parse("https://zenn.dev/")
		if err != nil {
			return fmt.Errorf("failed to parse zenn url: %w", err)
		}

		items, err := u.rssClient.FetchItems(ctx, zu.JoinPath(input.zennUserID, "feed"))
		if err != nil {
			return fmt.Errorf("failed to fetch zenn items: %w", err)
		}

		profile, err = profile.ReplaceZenn(ToMarkdown(items, input.zennMaxArticles))
		if err != nil {
			return fmt.Errorf("failed to replace zenn: %w", err)
		}
	}

	if input.connpassNickname != "" {
		profile, err = func(input UpdateProfileUsecaseInput, profile *Profile) (*Profile, error) {
			const readmeURL = "https://github.com/kumackey/profile-updater?tab=readme-ov-file#connpass"

			if time.Now().After(time.Date(2024, 5, 23, 0, 0, 0, 0, time.UTC)) {
				fmt.Printf("WARNING: connpassのサポートは2024年5月23日を以て廃止されました。\n"+
					"connpassの処理はスキップされます。\n"+
					"詳細はREADMEをご確認ください: %s\n", readmeURL,
				)
				return profile, fmt.Errorf("connpassのサポートは廃止しました。詳細はREADMEをご確認ください: %s", readmeURL)
			}

			fmt.Printf("WARNING: connpassのサポートは2024年5月23日以降に廃止されます。詳細はREADMEをご確認ください: %s\n", readmeURL)

			connpassList, err := u.connpassClient.FetchEventList(ctx, input.connpassNickname)
			if err != nil {
				return profile, err
			}

			replaceValue := ToMarkdown(connpassList, input.connpassMaxEvents)

			return profile.ReplaceConnpass(replaceValue)
		}(input, profile)
		if err != nil {
			return err
		}
	}

	if input.qiitaUserID != "" {
		qiitaArticleList, err := u.qiitaClient.FetchArticleList(ctx, input.qiitaUserID, input.qiitaMaxArticles)
		if err != nil {
			return err
		}

		replaceValue := ToMarkdown(qiitaArticleList, input.qiitaMaxArticles)

		profile, err = profile.ReplaceQiita(replaceValue)
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
	profileIO ProfileIO, connpassClient ConnpassClient, qiitaClient QiitaClient, RssClient RssClient,
) UpdateProfileUsecase {
	return UpdateProfileUsecase{
		profileIO:      profileIO,
		connpassClient: connpassClient,
		qiitaClient:    qiitaClient,
		rssClient:      RssClient,
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
