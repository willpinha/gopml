package opml_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
	"github.com/willpinha/gopml/opml"
)

func TestMain(m *testing.M) {
	v := m.Run()

	snaps.Clean(m)

	os.Exit(v)
}

func TestMarshalOutline(t *testing.T) {
	tests := []opml.Outline{
		{Text: "Hello world"},
		{},
	}

	for _, test := range tests {
		data, err := xml.Marshal(test)

		require.NoError(t, err)
		snaps.MatchStandaloneSnapshot(t, string(data))
	}
}
