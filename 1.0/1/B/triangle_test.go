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
			input: "3\n4\n5\n",
			want:  "YES",
		},
		{
			name:  "example test #2",
			input: "3\n5\n4\n",
			want:  "YES",
		},
		{
			name:  "example test #3",
			input: "4\n5\n3\n",
			want:  "YES",
		},
		{
			name:  "NO answer",
			input: "40\n5\n3\n",
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
