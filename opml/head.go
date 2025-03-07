package opml

import (
	"encoding/xml"
)

type Head struct {
	XMLName         xml.Name `xml:"head"`
	Title           string   `xml:"title"`
	DateCreated     OpmlTime `xml:"dateCreated"`
	DateModified    OpmlTime `xml:"dateModified"`
	OwnerName       string   `xml:"ownerName"`
	OwnerEmail      string   `xml:"ownerEmail"`
	OwnerId         string   `xml:"ownerId"`
	Docs            string   `xml:"docs"`
	ExpansionState  string   `xml:"expansionState"`
	VertScrollState int      `xml:"vertScrollState"`
	WindowTop       int      `xml:"windowTop"`
	WindowLeft      int      `xml:"windowLeft"`
	WindowBottom    int      `xml:"windowBottom"`
	WindowRight     int      `xml:"windowRight"`
}
