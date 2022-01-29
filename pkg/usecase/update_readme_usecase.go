package usecase

import (
	"github.com/kumackey/qiita-profile/pkg/domain"
)

type UpdateReadmeUsecase struct {
	ReadmeFile ReadmeFile
}

func (u UpdateReadmeUsecase) Exec() error {
	readme, err := u.ReadmeFile.Scan()
	if err != nil {
		return err
	}

	readme = readme.Replace([]domain.Line{"書き換えました"})
	err = u.ReadmeFile.Write(readme)
	if err != nil {
		return err
	}

	return nil
}
