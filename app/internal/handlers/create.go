package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/k0pernicus/go-tinyurl/internal/helpers"

	app "github.com/k0pernicus/go-tinyurl/internal"
	"github.com/k0pernicus/go-tinyurl/pkg/types"
	qrcode "github.com/skip2/go-qrcode"
)

// Create permits to create a tiny URL
func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
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

	hasDeadline := c.DeadIn.Duration != 0
	app.DB.Store(id, app.Informations{
		Redirection: c.URL,
		HasDeadline: hasDeadline,
		Deadline:    time.Now().Add(c.DeadIn.Duration),
	})

	response := types.CreationResponse{
		Message: types.OK,
		ID:      id,
	}

	if c.GenerateQRCode {
		if q, err := qrcode.New(c.URL, qrcode.Medium); err != nil {
			fmt.Printf("Failed to generate qr code: %s\n", err.Error())
		} else {
			response.QRCode = q.ToSmallString(false)
		}
	}

	helpers.AnswerWith(w, types.Response{
		StatusCode: http.StatusOK,
		Response:   response,
	})
}
