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
			input: "10 2 2 10",
			want:  "20 2\n2 20\n4 10\n10 4",
		},
		{
			name:  "example test #2",
			input: "5 7 3 2",
			want:  "9 5\n5 9",
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
