package helpers

import (
	"fmt"

	app "github.com/k0pernicus/go-tinyurl/internal"
)

func BuildURL(id string) string {
	return fmt.Sprintf("http://%s:%s/%s", app.C.Host, app.C.Port, id)
}
