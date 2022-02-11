package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_ReplaceZenn(t *testing.T) {
	tests := map[string]struct {
		input    *Profile
		replaces string
		output   *Profile
	}{
		"Zennの置き換えの記述が存在する": {
			input:    &Profile{Content: "こんにちは\n" + regexZennBegin + "\n書き換えられる前です\n" + regexZennEnd},
			replaces: "\n書き換えました\n",
			output: &Profile{
				Content: "こんにちは\n" + regexZennBegin + "\n書き換えました\n" + regexZennEnd,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile, err := test.input.ReplaceZenn(test.replaces)
			assert.Nil(t, err)
			assert.Equal(t, *test.output, *profile)
		})
	}
}

func TestProfile_ReplaceZenn_Failed(t *testing.T) {
	tests := map[string]struct {
		input *Profile
	}{
		"置き換えの記述が存在しない": {
			input: &Profile{Content: "こんにちは"},
		},
		"Connpassの置き換えの記述が存在している": {
			input: &Profile{Content: "こんにちは\n" + regexConnpassBegin + "\n書き換えられる前です\n" + regexConnpassEnd},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.input.ReplaceZenn("置き換えたい文章")
			assert.Equal(t, ErrReplaceStatementNotFound, err)
		})
	}
}

func TestProfile_ReplaceConnpass(t *testing.T) {
	tests := map[string]struct {
		input    *Profile
		replaces string
		output   *Profile
	}{
		"Connpassの置き換えの記述が存在する": {
			input:    &Profile{Content: "こんにちは\n" + regexConnpassBegin + "\n書き換えられる前です\n" + regexConnpassEnd},
			replaces: "\n書き換えました\n",
			output: &Profile{
				Content: "こんにちは\n" + regexConnpassBegin + "\n書き換えました\n" + regexConnpassEnd,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile, err := test.input.ReplaceConnpass(test.replaces)
			assert.Nil(t, err)
			assert.Equal(t, *test.output, *profile)
		})
	}
}

func TestProfile_ReplaceQiita(t *testing.T) {
	tests := map[string]struct {
		input    *Profile
		replaces string
		output   *Profile
	}{
		"qiitaの置き換えの記述が存在する": {
			input:    &Profile{Content: "こんにちは\n" + regexQiitaBegin + "\n書き換えられる前です\n" + regexQiitaEnd},
			replaces: "\n書き換えました\n",
			output: &Profile{
				Content: "こんにちは\n" + regexQiitaBegin + "\n書き換えました\n" + regexQiitaEnd,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile, err := test.input.ReplaceQiita(test.replaces)
			assert.Nil(t, err)
			assert.Equal(t, *test.output, *profile)
		})
	}
}
