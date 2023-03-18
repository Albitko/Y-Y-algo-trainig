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
			input: "89 20 41 1 11",
			want:  "2 3",
		},
		{
			name:  "example test #2",
			input: "11 1 1 1 1",
			want:  "0 1",
		},
		{
			name:  "example test #3",
			input: "3 2 2 2 1",
			want:  "-1 -1",
		},
		{
			name:  "first front door upper than given",
			input: "67 20 41 1 11",
			want:  "1 17",
		},
		{
			name:  "N front door after target",
			input: "89 20 216 3 14",
			want:  "2 3",
		},
		{
			name:  "Undefined floor",
			input: "4 2 5 1 2",
			want:  "1 0",
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
