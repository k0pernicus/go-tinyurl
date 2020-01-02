package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/k0pernicus/go-tinyurl/internal/db"

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
	if err := db.AddRecord(app.DB, id, c.URL, time.Now().Add(c.DeadIn.Duration), hasDeadline); err != nil {
		fmt.Printf("Error when adding record in DB: %s\n", err.Error())
		helpers.AnswerWith(w, types.Response{
			StatusCode: http.StatusInternalServerError,
			Response: types.CreationResponse{
				Message: types.CannotInsertRecord,
			},
		})
		return
	}

	if c.GenerateQRCode {
		if q, err := qrcode.New(helpers.BuildURL(id), qrcode.Medium); err != nil {
			fmt.Printf("Failed to generate qr code: %s\n", err.Error())
			helpers.AnswerWith(w, types.Response{
				StatusCode: http.StatusInternalServerError,
				Response: types.CreationResponse{
					Message: types.CannotCreateQRCode,
				},
			})
		} else {
			helpers.AnswerRaw(w, types.Response{
				StatusCode: http.StatusOK,
				Response:   q.ToSmallString(false),
			})
		}
		return
	}

	helpers.AnswerWith(w, types.Response{
		StatusCode: http.StatusOK,
		Response: types.CreationResponse{
			Message: types.OK,
			ID:      id,
		},
	})
}
