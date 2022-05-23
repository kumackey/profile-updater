package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	zennUserID := os.Getenv("INPUT_ZENN_USER_ID")
	zennMaxArticles, err := getMaxLines(os.Getenv("INPUT_ZENN_MAX_ARTICLES"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connpassNickname := os.Getenv("INPUT_CONNPASS_NICKNAME")
	connpassMaxEvents, err := getMaxLines(os.Getenv("INPUT_CONNPASS_MAX_EVENTS"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	qiitaUserID := os.Getenv("INPUT_QIITA_USER_ID")
	qiitaMaxArticles, err := getMaxLines(os.Getenv("INPUT_QIITA_MAX_ARTICLES"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if zennUserID == "" && connpassNickname == "" && qiitaUserID == "" {
		fmt.Println("zenn user id or connpass nickname required")
		os.Exit(1)
	}

	u := usecase.NewUpdateProfileUsecase(
		adapter.ReadmeFileOS{}, adapter.ZennRSSClient{}, adapter.ConnpassAPIClient{}, adapter.QiitaAPIClient{},
	)
	input := usecase.NewUpdateProfileUseCaseInput(
		zennUserID,
		zennMaxArticles,
		connpassNickname,
		connpassMaxEvents,
		qiitaUserID,
		qiitaMaxArticles,
	)

	err = u.Exec(context.Background(), input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}

func getMaxLines(v string) (int, error) {
	if v == "" {
		return usecase.DefaultMaxLines, nil
	}

	maxArticles, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return maxArticles, nil
}
