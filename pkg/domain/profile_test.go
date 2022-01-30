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
					"こんにちは",
					beginLine,
					"置き換えられるライン",
					endLine,
				}},
			replaces: []string{
				"書き換えました",
			},
			output: &Profile{
				Content: []line{
					"こんにちは",
					beginLine,
					"書き換えました",
					endLine,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, *test.output, *test.input.Replace(test.replaces))
		})
	}
}
