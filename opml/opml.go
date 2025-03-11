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

// Reads the data defined in the reader and converts it to Opml.
func ReadFrom(r io.Reader) (*Opml, error) {
	var o Opml

	err := xml.NewDecoder(r).Decode(&o)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

// Reads the data defined in a file and converts it to Opml.
func ReadFromFile(f *os.File) (*Opml, error) {
	return ReadFrom(f)
}

// Reads data returned from a URL (as an OPML document) and converts it to Opml.
func ReadFromURL(url string) (*Opml, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ReadFrom(resp.Body)
}

// Write Opml to a writer.
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

// Write Opml to a file.
func (o *Opml) WriteToFile(f *os.File) (int64, error) {
	return o.WriteTo(f)
}
