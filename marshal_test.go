package base_test

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
	base "github.com/willpinha/gopml"
)

func TestMain(m *testing.M) {
	v := m.Run()

	snaps.Clean(m)

	os.Exit(v)
}

func TestMarshal(t *testing.T) {
	o := base.Outline{
		Text: "Hello",
		Outlines: []base.Outline{
			{
				Text: "World",
			},
		},
	}

	data, err := o.Marshal()

	require.NoError(t, err)
	snaps.MatchSnapshot(t, string(data))
}
