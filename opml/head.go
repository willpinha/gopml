package opml

import (
	"encoding/xml"
)

type Head struct {
	XMLName         xml.Name  `xml:"head,omitempty"`
	Title           string    `xml:"title,omitempty"`
	DateCreated     *OpmlTime `xml:"dateCreated,omitempty"`
	DateModified    *OpmlTime `xml:"dateModified,omitempty"`
	OwnerName       string    `xml:"ownerName,omitempty"`
	OwnerEmail      string    `xml:"ownerEmail,omitempty"`
	OwnerId         string    `xml:"ownerId,omitempty"`
	Docs            string    `xml:"docs,omitempty"`
	ExpansionState  string    `xml:"expansionState,omitempty"`
	VertScrollState int       `xml:"vertScrollState,omitempty"`
	WindowTop       int       `xml:"windowTop,omitempty"`
	WindowLeft      int       `xml:"windowLeft,omitempty"`
	WindowBottom    int       `xml:"windowBottom,omitempty"`
	WindowRight     int       `xml:"windowRight,omitempty"`
}

func (h *Head) ExtractExpansionState() ([]int, error) {
	return nil, nil
}
