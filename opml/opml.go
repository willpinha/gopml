package opml

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
)

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Version Version  `xml:"version,attr"`
	Head    *Head    `xml:"head"`
	Body    *Body    `xml:"body"`
}

func LoadFromFile(filename string) (*Opml, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return &Opml{}, nil
	}

	var o *Opml

	err = xml.Unmarshal(data, o)

	if err != nil {
		return &Opml{}, nil
	}

	return o, nil
}

func LoadFromURL(url string) (*Opml, error) {
	resp, err := http.Get(url)

	if err != nil {
		return &Opml{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return &Opml{}, err
	}

	var o *Opml

	err = xml.Unmarshal(body, o)

	if err != nil {
		return &Opml{}, err
	}

	return o, nil
}

func (o *Opml) SaveToFile(filename string) error {
	data, err := xml.MarshalIndent(o, "\t", "\t\t")

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (o *Opml) ForEachExpansionState(callback func(outline Outline)) {}
