package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestTask(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "simple test #1",
			input: "6\n7 1 5 3 6 4\n",
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader = strings.NewReader(tt.input)
			result := problem(reader)
			fmt.Print(result, tt.want)
		})
	}

}
