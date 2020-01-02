package db

import (
	"time"
)

type Record struct {
	ID          string
	Redirection string
	Deadline    time.Time
	HasDeadline bool
}
