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

func (m *connpassClientMock) FetchEventList(ctx context.Context, userNickname string) (domain.ConpassEventList, error) {
	ret := m.Called(ctx, userNickname)

	return ret.Get(0).(domain.ConpassEventList), ret.Error(1)
}

type qiitaClientMock struct {
	mock.Mock
}

func (m *qiitaClientMock) FetchArticleList(ctx context.Context, userID string, limit int) (domain.QiitaArticleList, error) {
	ret := m.Called(ctx, userID, limit)

	return ret.Get(0).(domain.QiitaArticleList), ret.Error(1)
}

func TestUpdateProfileUsecase_Exec(t *testing.T) {
	var tests = map[string]struct {
		input            UpdateProfileUsecaseInput
		retProfileIOScan *domain.Profile
		output           error
	}{
		"全部の値が入っている": {
			input: UpdateProfileUsecaseInput{
				zennUserID:        "kumackey",
				zennMaxArticles:   5,
				connpassNickname:  "kumackey",
				connpassMaxEvents: 5,
				qiitaUserID:       "kumackey",
				qiitaMaxArticles:  5,
			},
			retProfileIOScan: &domain.Profile{
				Content: "<!-- profile updater begin: zenn --><!-- profile updater end: zenn -->" +
					"<!-- profile updater begin: connpass --><!-- profile updater end: connpass -->" +
					"<!-- profile updater begin: qiita --><!-- profile updater end: qiita -->",
			},
			output: nil,
		},
		"Zennだけの値が入っている": {
			input: UpdateProfileUsecaseInput{
				zennUserID:        "kumackey",
				zennMaxArticles:   5,
				connpassMaxEvents: 5,
				qiitaMaxArticles:  5,
			},
			retProfileIOScan: &domain.Profile{
				Content: "<!-- profile updater begin: zenn --><!-- profile updater end: zenn -->",
			},
			output: nil,
		},
		"Zennの値が入っているのに、プロフィールに該当する置換箇所がない": {
			input: UpdateProfileUsecaseInput{
				zennUserID:        "kumackey",
				zennMaxArticles:   5,
				connpassMaxEvents: 5,
				qiitaMaxArticles:  5,
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
			qiitaClientMock := new(qiitaClientMock)
			usecase := UpdateProfileUsecase{profileIOMock, zennClientMock, connpassClientMock, qiitaClientMock}

			profileIOMock.On("Write", mock.Anything).Return(nil)
			zennClientMock.On("FetchArticleList", mock.Anything, mock.Anything).Return(domain.ZennArticleList{}, nil)
			connpassClientMock.On("FetchEventList", mock.Anything, mock.Anything).Return(domain.ConpassEventList{}, nil)
			profileIOMock.On("Scan").Return(test.retProfileIOScan, nil)
			qiitaClientMock.On("FetchArticleList", mock.Anything, mock.Anything, mock.Anything).Return(domain.QiitaArticleList{}, nil)

			err := usecase.Exec(context.Background(), test.input)
			assert.Equal(t, test.output, err)
		})
	}
}
