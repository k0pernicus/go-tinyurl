package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	app "github.com/k0pernicus/go-tinyurl/internal"
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

	object, exists := app.DB.Load(id)
	if !exists {
		fmt.Println("ID does not exists")
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusNotFound,
			Response: types.ExistsResponse{
				Message: types.OK,
			},
		})
		return
	}

	informations := object.(app.Informations)

	if informations.IsDead() {
		app.DB.Delete(id)
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusNotFound,
			Response: types.ExistsResponse{
				Message: types.OK,
			},
		})
		return
	}

	http.Redirect(w, r, informations.Redirection, http.StatusFound)
}
