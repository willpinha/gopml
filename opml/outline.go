package opml

import (
	"encoding/xml"
)

type Body struct {
	XMLName  xml.Name  `xml:"body"`
	Outlines []Outline `xml:"outline"`
}

type Outline struct {
	XMLName      xml.Name  `xml:"outline"`
	Text         string    `xml:"text,attr"`
	Type         string    `xml:"type,attr"`
	IsCommented  bool      `xml:"isCommented,attr"`
	IsBreakpoint bool      `xml:"isBreakpoint,attr"`
	Created      OpmlTime  `xml:"created,attr"`
	Category     string    `xml:"category,attr"`
	Outlines     []Outline `xml:"outline"`
}
