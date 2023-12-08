package bexio

import (
	"strings"
	"time"
)

type Timestamp struct {
	time.Time
}

const timestampLayout = "2006-01-02 15:04:05"

func (t *Timestamp) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(timestampLayout, s)
	return
}
