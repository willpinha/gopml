package opml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willpinha/gopml/opml"
)

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		input  opml.Version
		output error
	}{
		{input: "invalid", output: opml.ErrUnknownPomlVersion},
		{input: "1.0", output: opml.ErrUnsupportedPomlVersion},
		{input: "1.1", output: opml.ErrUnsupportedPomlVersion},
		{input: "2.0", output: nil},
	}

	for _, test := range tests {
		o := opml.Opml{Version: test.input}

		err := o.ValidateVersion()

		assert.ErrorIs(t, err, test.output)
	}
}
