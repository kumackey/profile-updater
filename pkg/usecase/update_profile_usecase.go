package usecase

import (
	"fmt"
)

type UpdateProfileUsecase struct {
	ProfileIO ProfileIO
}

func (u UpdateProfileUsecase) Exec() error {
	readme, err := u.ProfileIO.Scan()
	if err != nil {
		return err
	}

	readme, err = readme.Replace([]string{"書き換えました"})
	if err != nil {
		return err
	}

	err = u.ProfileIO.Write(readme)
	if err != nil {
		return err
	}

	fmt.Println(*readme)

	return nil
}
