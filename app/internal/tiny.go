package tiny

import (
	"database/sql"
	"fmt"
	"time"
)

var (
	// C is the internal service configuration variable
	C Configuration
	// DB is a temporary internal database that contains all created tiny URL and redirections
	DB *sql.DB
)

func ConnectDB(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	_, err = DB.Exec("create table if not exists urls (id text, redirection text, deadline time, has_deadline bool)")
	return err
}

// Configuration handles all the informations for creation & launch
type Configuration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// String returns, as a string, the app as HOST:PORT
func (c Configuration) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

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
