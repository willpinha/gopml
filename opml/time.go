package opml

import "time"

type OpmlTime time.Time

func (t OpmlTime) String() string {
	return time.Time(t).Format(time.RFC822)
}

func ParseOpmlTime(value string) (OpmlTime, error) {
	t, err := time.Parse(time.RFC822, value)

	if err != nil {
		return OpmlTime{}, err
	}

	return OpmlTime(t), nil
}
