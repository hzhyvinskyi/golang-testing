package cases

import "time"

type Payload struct {
	ID        uint64
	Command   string
	Title     string
	Timestamp time.Time
}

type Observer interface {
	getID() string
	update(*Payload)
}
