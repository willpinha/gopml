package opml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Head struct {
	XMLName         xml.Name  `xml:"head"`
	Title           string    `xml:"title,omitempty"`
	DateCreated     *OpmlTime `xml:"dateCreated,omitempty"`
	DateModified    *OpmlTime `xml:"dateModified,omitempty"`
	OwnerName       string    `xml:"ownerName,omitempty"`
	OwnerEmail      string    `xml:"ownerEmail,omitempty"`
	OwnerId         string    `xml:"ownerId,omitempty"`
	Docs            string    `xml:"docs,omitempty"`
	ExpansionState  string    `xml:"expansionState,omitempty"`
	VertScrollState int       `xml:"vertScrollState,omitempty"`
	WindowTop       int       `xml:"windowTop,omitempty"`
	WindowLeft      int       `xml:"windowLeft,omitempty"`
	WindowBottom    int       `xml:"windowBottom,omitempty"`
	WindowRight     int       `xml:"windowRight,omitempty"`
}

func (h *Head) ExtractExpansionState() ([]int, error) {
	if h.ExpansionState == "" {
		return []int{}, nil
	}

	split := strings.Split(h.ExpansionState, ",")

	numbers := make([]int, len(split))

	for i, s := range split {
		sTrim := strings.TrimSpace(s)

		n, err := strconv.Atoi(sTrim)

		if err != nil {
			return nil, fmt.Errorf("could not extract expansion state: %v", err)
		}

		numbers[i] = n
	}

	return numbers, nil
}
