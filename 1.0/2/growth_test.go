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
		want  string
	}{
		{
			name:  "example test #1",
			input: "1 7 9",
			want:  "YES",
		},
		{
			name:  "example test #2",
			input: "1 9 7",
			want:  "NO",
		},
		{
			name:  "example test #3",
			input: "2 2 2",
			want:  "NO",
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
