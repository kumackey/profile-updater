package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kumackey/profile-updater/internal/adapter"
	"github.com/kumackey/profile-updater/internal/domain"
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
	qiitaSortByLgtm := getBool(os.Getenv("INPUT_QIITA_SORT_BY_LGTM"))

	if zennUserID == "" && connpassNickname == "" && qiitaUserID == "" {
		fmt.Println("zenn user id or connpass nickname required")
		os.Exit(1)
	}

	u := domain.NewUpdateProfileUsecase(
		adapter.ReadmeFileOS{},
		adapter.ConnpassAPIClient{},
		adapter.QiitaAPIClient{},
		adapter.NewRssClient(http.DefaultClient),
	)
	input := domain.NewUpdateProfileUseCaseInput(
		zennUserID,
		zennMaxArticles,
		connpassNickname,
		connpassMaxEvents,
		qiitaUserID,
		qiitaMaxArticles,
		qiitaSortByLgtm,
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
		return domain.DefaultMaxLines, nil
	}

	maxArticles, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return maxArticles, nil
}

func getBool(v string) bool {
	if v == "true" || v == "1" {
		return true
	}
	return false
}
