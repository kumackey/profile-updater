package main

import (
	"fmt"

	"github.com/kumackey/profile-updater/pkg/adapter"
	"github.com/kumackey/profile-updater/pkg/usecase"
)

func main() {
	u := usecase.UpdateProfileUsecase{ProfileIO: adapter.ReadmeFileOS{}}
	err := u.Exec()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success")
}
