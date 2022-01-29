package main

import (
	"fmt"

	"github.com/kumackey/qiita-profile/pkg/usecase"
)

func main() {
	u := usecase.NewUpdateReadMe()
	err := u.Exec()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success")
}
