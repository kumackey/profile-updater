package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	zennUserID := os.Getenv("INPUT_ZENN_USER_ID")
	if zennUserID == "" {
		fmt.Println("zenn user id required")
		os.Exit(1)
	}

	zennMaxArticles, err := getZennMaxArticles(os.Getenv("INPUT_ZENN_MAX_ARTICLES"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	u := usecase.NewUpdateProfileUsecase(adapter.ReadmeFileOS{}, adapter.ZennRSSClient{})
	err = u.Exec(context.Background(), zennUserID, zennMaxArticles)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}

func getZennMaxArticles(v string) (int, error) {
	if v == "" {
		return domain.DefaultZennMaxArticles, nil
	}

	maxArticles, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return maxArticles, nil
}
