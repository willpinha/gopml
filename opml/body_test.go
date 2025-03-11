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

func TestMarshalOutline(t *testing.T) {
	mockOpmlTime := opml.OpmlTime(time.Unix(0, 0).UTC())

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

func TestCustomOutline(t *testing.T) {
	type CustomOutline struct {
		opml.Outline
		Hello string `xml:"hello,attr"`
		World int    `xml:"world"`
	}

	c := CustomOutline{
		Outline: opml.Outline{
			Text: "Goodbye everybody",
		},
		Hello: "hello",
		World: 1,
	}

	data, err := xml.Marshal(c)

	require.NoError(t, err)
	snaps.MatchSnapshot(t, string(data))
}
