package encoding

import "encoding/xml"

type Outline struct {
	XMLName xml.Name `xml:"outline"`
	Text    string   `xml:"text,attr"`
}
