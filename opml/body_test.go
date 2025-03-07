package opml_test

import (
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
	"github.com/willpinha/gopml/opml"
)

func TestMain(m *testing.M) {
	v := m.Run()

	snaps.Clean(m)

	os.Exit(v)
}

var (
	mockOpmlTime = opml.OpmlTime(time.Unix(0, 0).UTC())
)

func TestMarshalOutline(t *testing.T) {
	tests := []opml.Outline{
		{},
		{
			Text: "Hello world",
		},
		{
			Type:         "hello",
			Text:         "Hello world",
			IsCommented:  true,
			IsBreakpoint: false,
			Created:      &mockOpmlTime,
			Category:     "world",
		},
		{
			Text: "Names",
			Outlines: []*opml.Outline{
				{Text: "Will"},
				{Text: "Bob"},
				{Text: "Jack"},
			},
		},
	}

	for _, test := range tests {
		data, err := xml.MarshalIndent(test, "", "\t")

		require.NoError(t, err)
		snaps.MatchSnapshot(t, string(data))
	}
}

func TestMarshalBody(t *testing.T) {
	b := opml.Body{
		Outlines: []*opml.Outline{
			{
				Text: "Hello",
				Outlines: []*opml.Outline{
					{Text: "World"},
				},
			},
			{},
		},
	}

	data, err := xml.MarshalIndent(b, "", "\t")

	require.NoError(t, err)
	snaps.MatchSnapshot(t, string(data))
}
