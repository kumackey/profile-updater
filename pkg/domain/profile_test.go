package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_Replace(t *testing.T) {
	tests := map[string]struct {
		input    *Profile
		replaces []string
		output   *Profile
	}{
		"beginとendの両方がある": {
			input: &Profile{
				Content: []line{
					{"こんにちは", false},
					{"", true},
				},
			},
			replaces: []string{
				"書き換えました",
			},
			output: &Profile{
				Content: []line{
					{"こんにちは", false},
					{beginLine, false},
					{"書き換えました", false},
					{endLine, false},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			profile, err := test.input.Replace(test.replaces)
			assert.Nil(t, err)
			assert.Equal(t, *test.output, *profile)
		})
	}
}

func TestProfile_Replace_Failed(t *testing.T) {
	tests := map[string]struct {
		input *Profile
	}{
		"置き換えフラグ箇所が存在しない": {
			input: &Profile{
				Content: []line{
					{"こんにちは", false},
				}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.input.Replace([]string{"置き換えた後の文章"})
			assert.Equal(t, ErrReplaceLinesNotFound, err)
		})
	}
}
