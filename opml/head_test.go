package opml_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willpinha/gopml/opml"
)

func TestExtractExpansionState(t *testing.T) {
	tests := []struct {
		input  string
		output []int // nil if error expected
	}{
		{input: "1,2,3,4", output: []int{1, 2, 3, 4}},
		{input: "", output: []int{}},
		{input: "  1  ,  2  ,  3  ,  4  ", output: []int{1, 2, 3, 4}},
		{input: ",1,2,", output: nil},
		{input: "1,a,2", output: nil},
	}

	for _, test := range tests {
		head := &opml.Head{ExpansionState: test.input}

		es, err := head.ExtractExpansionState()

		if test.output == nil {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, test.output, es)
		}
	}
}
