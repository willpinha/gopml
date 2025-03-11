package opml

import (
	"encoding/xml"
)

type Body struct {
	XMLName  xml.Name   `xml:"body"`
	Outlines []*Outline `xml:"outline"`
}

// An outline is an atomic unit of information in an OPML document. This struct contains only
// the fields defined in the OPML 2.0 specification, but it is possible to extend it using
// composition:
//
//	type CustomOutline struct {
//		opml.Outline
//		Hello string `xml:"hello,attr"`
//		World bool   `xml:"world,attr"`
//	}
//
// It is from composition that the outlines of the "link" and "include" types are implemented.
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

// A custom outline that represents a link to a URL.
type LinkOutline struct {
	Outline
	URL string `xml:"url,attr"`
}

// A custom outline that represents a link to a URL that returns an OPML document (text/xml+opml).
type IncludeOutline struct {
	Outline
	URL string `xml:"url,attr"`
}

// Overrides the child outlines with those defined in the URL's OPML document.
func (o *IncludeOutline) ReadOutlines() error {
	externalOpml, err := ReadFromURL(o.URL)

	if err != nil {
		return err
	}

	o.Outlines = externalOpml.Body.Outlines

	return nil
}

// A custom outline for subscription lists with RSS.
type RSSOutline struct {
	Outline
	XmlURL      string `xml:"xmlUrl,attr"`
	Description string `xml:"description,attr,omitempty"`
	HtmlURL     string `xml:"htmlUrl,attr,omitempty"`
	Language    string `xml:"language,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Version     string `xml:"version,attr,omitempty"`
}
