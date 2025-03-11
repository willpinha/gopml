package opml

import (
	"errors"
	"slices"
)

type Version string

const (
	Version1   Version = "1.0"
	Version1_1 Version = "1.1"
	Version2   Version = "2.0"
)

var (
	ErrUnknownOpmlVersion     = errors.New("unknown OPML version (only versions 1.0, 1.1 and 2.0 are valid)")
	ErrUnsupportedOpmlVersion = errors.New("unsupported OPML version (gopml only supports version 2.0)")
)

// Validates the version of the OPML document. The only accepted version for this package is "2.0".
func (o *Opml) ValidateVersion() error {
	validVersions := []Version{Version1, Version1_1, Version2}

	if !slices.Contains(validVersions, o.Version) {
		return ErrUnknownOpmlVersion
	}

	if o.Version != Version2 {
		return ErrUnsupportedOpmlVersion
	}

	return nil
}
