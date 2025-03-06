package base

import "encoding/xml"

type Outline struct {
	XMLName  xml.Name  `xml:"outline"`
	Text     string    `xml:"text,attr"`
	Outlines []Outline `xml:"outline"`
}

func (o *Outline) Marshal() ([]byte, error) {
	return xml.MarshalIndent(o, "\t", "\t")
}

func (o *Outline) Unmarshal() {

}
