package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/k0pernicus/go-tinyurl/internal/helpers"

	app "github.com/k0pernicus/go-tinyurl/internal"
	"github.com/k0pernicus/go-tinyurl/pkg/types"
)

// Create permits to create a tiny URL
func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c types.CreationRequest
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Printf("Error when decoding message: %s", err.Error())
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusBadRequest,
			Response: types.CreationResponse{
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	id := helpers.Generate()

	hasDeadline := c.Time != 0
	app.DB.Store(id, app.Informations{
		Redirection: c.URL,
		HasDeadline: hasDeadline,
		Deadline:    time.Now().Add(c.Time),
	})

	helpers.AnswerWith(w, types.Response{
		StatusCode: http.StatusOK,
		Response: types.CreationResponse{
			Message: types.OK,
			ID:      id,
		},
	})
}
