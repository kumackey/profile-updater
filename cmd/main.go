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

	if zennUserID == "" && connpassNickname == "" {
		fmt.Println("zenn user id or connpass nickname required")
		os.Exit(1)
	}

	u := usecase.NewUpdateProfileUsecase(adapter.ReadmeFileOS{}, adapter.ZennRSSClient{}, adapter.ConnpassAPIClient{})
	err = u.Exec(context.Background(), zennUserID, zennMaxArticles, connpassNickname, connpassMaxEvents)
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
