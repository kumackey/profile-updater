package main

import (
	"fmt"
	"os"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	u := usecase.UpdateProfileUsecase{ProfileIO: adapter.ReadmeFileOS{}, ZennClient: adapter.ZennRSS{}}
	err := u.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}
