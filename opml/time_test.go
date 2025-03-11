package opml_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/willpinha/gopml/opml"
)

func TestMarshalAndParseOpmlTime(t *testing.T) {
	mockOpmlTime := opml.OpmlTime(time.Unix(0, 0).UTC())
	s := "01 Jan 70 00:00 UTC"

	m, err := mockOpmlTime.MarshalText()

	require.NoError(t, err)
	require.Equal(t, s, string(m))

	p, err := opml.ParseOpmlTime(s)

	require.NoError(t, err)
	require.Equal(t, mockOpmlTime, p)
}
