package opml

import (
	"encoding/xml"
)

type Body struct {
	XMLName  xml.Name   `xml:"body"`
	Outlines []*Outline `xml:"outline"`
}

type Outline struct {
	XMLName      xml.Name   `xml:"outline"`
	Text         string     `xml:"text,attr"`
	Type         string     `xml:"type,attr,omitempty"`
	IsCommented  bool       `xml:"isCommented,attr,omitempty"`
	IsBreakpoint bool       `xml:"isBreakpoint,attr,omitempty"`
	Created      *OpmlTime  `xml:"created,attr,omitempty"`
	Category     string     `xml:"category,attr,omitempty"`
	Outlines     []*Outline `xml:"outline"`
}

type LinkOutline struct {
	Outline
	URL string `xml:"url,attr"`
}

type IncludeOutline struct {
	Outline
	URL string `xml:"url,attr"`
}

func (o *IncludeOutline) ReadOutlines() error {
	externalOpml, err := ReadFromURL(o.URL)

	if err != nil {
		return err
	}

	o.Outlines = externalOpml.Body.Outlines

	return nil
}

type RSSOutline struct {
	Outline
	XmlURL      string `xml:"xmlUrl,attr"`
	Description string `xml:"description,attr,omitempty"`
	HtmlURL     string `xml:"htmlUrl,attr,omitempty"`
	Language    string `xml:"language,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Version     string `xml:"version,attr,omitempty"`
}
