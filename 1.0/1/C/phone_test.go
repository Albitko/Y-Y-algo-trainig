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
			input: "8(495)430-23-97\n+7-4-9-5-43-023-97\n4-3-0-2-3-9-7\n8-495-430",
			want:  "YES\nYES\nNO\n",
		},
		{
			name:  "example test #2",
			input: "86406361642\n83341994118\n86406361642\n83341994118",
			want:  "NO\nYES\nNO\n",
		},
		{
			name:  "example test #3",
			input: "+78047952807\n+78047952807\n+76147514928\n88047952807",
			want:  "YES\nNO\nYES\n",
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
