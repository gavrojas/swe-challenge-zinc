package shared

import "time"

type Session struct {
	Uid        string
	ExpiryTime time.Time
}

var Sessions = make(map[string]Session)
