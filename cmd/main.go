package main

import (
	"fmt"
	"github.com/kumackey/qiita-profile/pkg/adapter"
	"github.com/kumackey/qiita-profile/pkg/usecase"
)

func main() {
	u := usecase.UpdateReadmeUsecase{ReadmeFile: adapter.ReadmeFileOS{}}
	err := u.Exec()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success")
}
