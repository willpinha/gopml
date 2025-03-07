package opml

import (
	"encoding/xml"
	"time"
)

type Head struct {
	XMLName         xml.Name  `xml:"head"`
	Title           string    `xml:"title"`
	DateCreated     time.Time `xml:"dateCreated"`
	DateModified    time.Time `xml:"dateModified"`
	OwnerName       string    `xml:"ownerName"`
	OwnerEmail      string    `xml:"ownerEmail"`
	OwnerId         string    `xml:"ownerId"`
	Docs            string    `xml:"docs"`
	ExpansionState  string    `xml:"expansionState"`
	VertScrollState int       `xml:"vertScrollState"`
	WindowTop       int       `xml:"windowTop"`
	WindowLeft      int       `xml:"windowLeft"`
	WindowBottom    int       `xml:"windowBottom"`
	WindowRight     int       `xml:"windowRight"`
}
