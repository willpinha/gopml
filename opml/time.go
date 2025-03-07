package opml

import (
	"time"
)

type OpmlTime time.Time

func (t OpmlTime) String() string {
	return t.ToTime().Format(time.RFC822)
}

func (t OpmlTime) ToTime() time.Time {
	return time.Time(t)
}

func ParseOpmlTime(value string) (OpmlTime, error) {
	t, err := time.Parse(time.RFC822, value)

	if err != nil {
		return OpmlTime{}, err
	}

	return OpmlTime(t), nil
}

func (t OpmlTime) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}
