package domain

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type profileIOMock struct {
	mock.Mock
}

func (m *profileIOMock) Scan() (*Profile, error) {
	ret := m.Called()

	return ret.Get(0).(*Profile), ret.Error(1)
}

func (m *profileIOMock) Write(*Profile) error {
	ret := m.Called()

	return ret.Error(0)
}

type connpassClientMock struct {
	mock.Mock
}

func (m *connpassClientMock) FetchEventList(ctx context.Context, userNickname string) ([]ConnpassEvent, error) {
	ret := m.Called(ctx, userNickname)

	return ret.Get(0).([]ConnpassEvent), ret.Error(1)
}

type qiitaClientMock struct {
	mock.Mock
}

func (m *qiitaClientMock) FetchArticleList(
	ctx context.Context,
	userID string,
	limit int,
) ([]QiitaArticle, error) {
	ret := m.Called(ctx, userID, limit)

	return ret.Get(0).([]QiitaArticle), ret.Error(1)
}

type rssClientMock struct {
	mock.Mock
}

func (m *rssClientMock) FetchItems(ctx context.Context, url *url.URL) ([]RssItem, error) {
	ret := m.Called(ctx, url)

	return ret.Get(0).([]RssItem), ret.Error(1)
}

func TestUpdateProfileUsecase_Exec(t *testing.T) {
	var tests = map[string]struct {
		input            UpdateProfileUsecaseInput
		retProfileIOScan *Profile
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
				// TODO: RSSのテストを追加する
			},
			retProfileIOScan: &Profile{
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
			retProfileIOScan: &Profile{
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
			retProfileIOScan: &Profile{},
			output:           ErrReplaceStatementNotFound,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profileIOMock := new(profileIOMock)
			connpassClientMock := new(connpassClientMock)
			qiitaClientMock := new(qiitaClientMock)
			rssClientMock := new(rssClientMock)
			usecase := UpdateProfileUsecase{profileIOMock, connpassClientMock, qiitaClientMock, rssClientMock}
			profileIOMock.On("Write", mock.Anything).Return(nil)
			connpassClientMock.On("FetchEventList", mock.Anything, mock.Anything).
				Return([]ConnpassEvent{}, nil)
			profileIOMock.On("Scan").
				Return(test.retProfileIOScan, nil)
			qiitaClientMock.On("FetchArticleList", mock.Anything, mock.Anything, mock.Anything).
				Return([]QiitaArticle{}, nil)
			rssClientMock.On("FetchItems", mock.Anything, mock.Anything).
				Return([]RssItem{}, nil)

			err := usecase.Exec(context.Background(), test.input)
			assert.Equal(t, test.output, err)
		})
	}
}
