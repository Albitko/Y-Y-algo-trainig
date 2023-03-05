package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestProblem(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example test #1",
			input: "10 20\nheat\n",
			want:  20,
		},
		{
			name:  "example test #2",
			input: "10 20\nfreeze\n",
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader = strings.NewReader(tt.input)
			result := problem(reader)
			assert.Equal(t, tt.want, result)
		})
	}

}
