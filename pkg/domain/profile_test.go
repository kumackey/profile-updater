package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_Replace(t *testing.T) {
	inputProfile := &Profile{
		Content: []Line{
			"こんにちは",
			beginLine,
			"置き換えられるライン",
			endLine,
		},
	}

	outputProfile := &Profile{
		Content: []Line{
			"こんにちは",
			beginLine,
			"書き換えました",
			endLine,
		},
	}

	tests := map[string]struct {
		input  *Profile
		output *Profile
	}{
		"beginとendの両方がある": {
			input:  inputProfile,
			output: outputProfile,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, *test.output, *test.input.Replace([]Line{"書き換えました"}))
		})
	}
}
