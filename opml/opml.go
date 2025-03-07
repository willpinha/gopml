package opml

import (
	"encoding/xml"
)

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Version Version  `xml:"version,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

func (o *Opml) IterateExpansionState(callback func(outline Outline)) {}
