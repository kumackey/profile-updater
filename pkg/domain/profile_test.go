package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_Replace(t *testing.T) {
	tests := map[string]struct {
		input    *Profile
		replaces string
		output   *Profile
	}{
		"beginとendの両方がある": {
			input:    &Profile{Content: "こんにちは\n" + regexBegin + "\n書き換えられる前です\n" + regexEnd},
			replaces: "\n書き換えました\n",
			output: &Profile{
				Content: "こんにちは\n" + regexBegin + "\n書き換えました\n" + regexEnd,
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
			input: &Profile{Content: "こんにちは"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.input.Replace("置き換えたい文章")
			assert.Equal(t, ErrReplaceLinesNotFound, err)
		})
	}
}
