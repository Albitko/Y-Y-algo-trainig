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
			input: "1\n3\n3\n2",
			want:  "5 7",
		},
		{
			name:  "example test #2",
			input: "1\n5\n1\n2",
			want:  "-1",
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
