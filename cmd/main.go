package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	u := usecase.NewUpdateProfileUsecase(adapter.ReadmeFileOS{}, adapter.ZennRSS{})
	err := u.Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}
