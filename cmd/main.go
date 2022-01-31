package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	zennUserID := os.Getenv("INPUT_ZENN_USER_ID")

	if zennUserID == "" {
		fmt.Println("zenn user id required")
		os.Exit(1)
	}

	u := usecase.NewUpdateProfileUsecase(adapter.ReadmeFileOS{}, adapter.ZennRSSClient{})
	err := u.Exec(context.Background(), zennUserID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}
