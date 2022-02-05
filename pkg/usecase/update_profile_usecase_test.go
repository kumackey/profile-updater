package usecase

import (
	"context"
	"testing"

	"github.com/kumackey/profile-updater/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type profileIOMock struct {
	mock.Mock
}

func (m *profileIOMock) Scan() (*domain.Profile, error) {
	ret := m.Called()

	return ret.Get(0).(*domain.Profile), ret.Error(1)
}

func (m *profileIOMock) Write(*domain.Profile) error {
	ret := m.Called()

	return ret.Error(0)
}

type zennClientMock struct {
	mock.Mock
}

func (m *zennClientMock) FetchArticleList(ctx context.Context, userID string) (domain.ZennArticleList, error) {
	ret := m.Called(ctx, userID)

	return ret.Get(0).(domain.ZennArticleList), ret.Error(1)
}

type connpassClientMock struct {
	mock.Mock
}

func (m *connpassClientMock) FetchEventList(ctx context.Context, userNickName string) (domain.ConpassEventList, error) {
	ret := m.Called(ctx, userNickName)

	return ret.Get(0).(domain.ConpassEventList), ret.Error(1)
}

func TestUpdateProfileUsecase_Exec(t *testing.T) {

	type input struct {
		zennUserID        string
		zennMaxArticles   int
		connpassNickName  string
		connpassMaxEvents int
	}

	tests := map[string]struct {
		input            input
		retProfileIOScan *domain.Profile
		output           error
	}{
		"ZennとConnpassの両方の値が入っている": {
			input: input{
				"kumackey", 5, "kumackey", 5,
			},
			retProfileIOScan: &domain.Profile{
				Content: "<!-- profile updater begin: zenn --><!-- profile updater end: zenn --><!-- profile updater begin: connpass --><!-- profile updater end: connpass -->",
			},
			output: nil,
		},
		"Zennだけの値が入っている": {
			input: input{
				"kumackey", 5, "", 5,
			},
			retProfileIOScan: &domain.Profile{
				Content: "<!-- profile updater begin: zenn --><!-- profile updater end: zenn -->",
			},
			output: nil,
		},
		"Zennの値が入っているのに、プロフィールに該当する置換箇所がない": {
			input: input{
				"kumackey", 5, "", 5,
			},
			retProfileIOScan: &domain.Profile{},
			output:           domain.ErrReplaceStatementNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profileIOMock := new(profileIOMock)
			zennClientMock := new(zennClientMock)
			connpassClientMock := new(connpassClientMock)
			usecase := UpdateProfileUsecase{profileIOMock, zennClientMock, connpassClientMock}

			profileIOMock.On("Write", mock.Anything).Return(nil)
			zennClientMock.On("FetchArticleList", mock.Anything, mock.Anything).Return(domain.ZennArticleList{}, nil)
			connpassClientMock.On("FetchEventList", mock.Anything, mock.Anything).Return(domain.ConpassEventList{}, nil)
			profileIOMock.On("Scan").Return(test.retProfileIOScan, nil)

			err := usecase.Exec(context.Background(), test.input.zennUserID, test.input.zennMaxArticles, test.input.connpassNickName, test.input.connpassMaxEvents)
			assert.Equal(t, test.output, err)
		})
	}
}
