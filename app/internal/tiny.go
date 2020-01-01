package tiny

import (
	"sync"
	"time"
)

// DB is a temporary internal database that contains all created tiny URL and redirections
var DB sync.Map

// Informations store tiny object informations
type Informations struct {
	Redirection string    `json:"redirection"`
	Deadline    time.Time `json:"deadline,omitempty"`
	HasDeadline bool      `json:"has_deadline"`
}

// IsDead permits to know if the object is dead now or not
func (t Informations) IsDead() bool {
	if !t.HasDeadline {
		return false
	}
	now := time.Now()
	return t.Deadline.Before(now)
}
