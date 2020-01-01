package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/k0pernicus/go-tinyurl/internal/helpers"

	"github.com/google/uuid"
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

	newUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Cannot create new UUID: %s", err.Error())
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusInternalServerError,
			Response: types.CreationResponse{
				Message: types.CannotGenerateNewUUID,
			},
		})
		return
	}

	hasDeadline := c.Time != 0
	app.DB.Store(newUUID, app.Informations{
		Redirection: c.URL,
		HasDeadline: hasDeadline,
		Deadline:    time.Now().Add(c.Time),
	})

	helpers.AnswerWith(w, types.Response{
		StatusCode: http.StatusOK,
		Response: types.CreationResponse{
			Message: types.OK,
			ID:      newUUID.String(),
		},
	})
}
