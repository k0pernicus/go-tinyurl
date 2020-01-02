package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	app "github.com/k0pernicus/go-tinyurl/internal"
	"github.com/k0pernicus/go-tinyurl/internal/db"
	"github.com/k0pernicus/go-tinyurl/internal/helpers"
	"github.com/k0pernicus/go-tinyurl/pkg/types"
)

// Exists permits to return if an ID has been found, and what is the redirected URL
func Exists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		fmt.Println("Cannot find 'id' query parameter in user's request")
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusBadRequest,
			Response: types.ExistsResponse{
				Message: types.CannotDecodeMessage,
			},
		})
		return
	}

	_, exists := db.GetRecord(app.DB, id)
	statusCode := http.StatusOK
	message := types.OK
	if exists != nil {
		statusCode = http.StatusNotFound
		message = types.URLDoesNotExists
	}
	helpers.AnswerWith(w, types.Response{
		StatusCode: statusCode,
		Response: types.ExistsResponse{
			Message: message,
			Exists:  exists == nil,
		},
	})
}
