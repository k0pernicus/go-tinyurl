package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	app "github.com/k0pernicus/go-tinyurl/internal"
	"github.com/k0pernicus/go-tinyurl/internal/db"
	"github.com/k0pernicus/go-tinyurl/internal/helpers"
	"github.com/k0pernicus/go-tinyurl/pkg/types"
)

// Get redirects to the right URL
func Get(w http.ResponseWriter, r *http.Request) {
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

	informations, exists := db.GetRecord(app.DB, id)
	if exists != nil {
		fmt.Printf("Failed to retrieve ID %s in DB\n", id)
		if exists == errors.New("Not found") {
			helpers.AnswerWith(w, types.Response{
				StatusCode: http.StatusNotFound,
				Response: types.ExistsResponse{
					Message: types.OK,
				},
			})
		} else {
			helpers.AnswerWith(w, types.Response{
				StatusCode: http.StatusInternalServerError,
				Response: types.ExistsResponse{
					Message: types.CannotRetrieveRecord,
				},
			})
		}
		return
	}

	if informations.HasDeadline && time.Now().After(informations.Deadline) {
		db.DeleteRecord(app.DB, id)
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusNotFound,
			Response: types.ExistsResponse{
				Message: types.URLDoesNotExists,
			},
		})
		return
	}

	http.Redirect(w, r, informations.Redirection, http.StatusFound)
}
