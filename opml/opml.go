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

func ReadFrom(r io.Reader) (*Opml, error) {
	var o *Opml

	err := xml.NewDecoder(r).Decode(o)

	if err != nil {
		return nil, err
	}

	return o, nil
}

func ReadFromFile(f *os.File) (*Opml, error) {
	return ReadFrom(f)
}

func ReadFromURL(url string) (*Opml, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ReadFrom(resp.Body)
}

func (o *Opml) WriteTo(w io.Writer) (int64, error) {
	data, err := xml.Marshal(o)

	if err != nil {
		return 0, err
	}

	n, err := w.Write(data)

	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (o *Opml) WriteToFile(f *os.File) (int64, error) {
	return o.WriteTo(f)
}
