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
	IsCommented  bool      `xml:"isCommented,attr,omitempty"`
	IsBreakpoint bool      `xml:"isBreakpoint,attr,omitempty"`
	Created      OpmlTime  `xml:"created,attr,omitempty"`
	Category     string    `xml:"category,attr,omitempty"`
	Outlines     []Outline `xml:"outline"`
}
