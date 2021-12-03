package utils

import (
	"errors"
	"time"
)

const (
	DATETIME_LAYOUT      = "2006-01-02 15:04:05"
	DATE_LAYOUT          = "2006-01-02"
	DATE_LAYOUT_SHORT_CN = "01月02日"
	DATE_LAYOUT_SHORT_EN = "01-02"
	TIME_LAYOUT          = "15:04:05"
	TIME_LAYOUT_SHORT    = "15:04"
)

type MyTime struct {
	time.Time
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == "\"\"" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	str := string(data)
	tt, err := time.ParseInLocation(`"`+DATETIME_LAYOUT+`"`, str, time.Local)
	t.Time = tt
	return err
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t *MyTime) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0)
	b = append(b, '"')
	b = t.AppendFormat(b, DATETIME_LAYOUT)
	b = append(b, '"')
	return b, nil
}
