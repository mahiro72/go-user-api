package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/mahiro72/go-user-api/pkg/logger"
)

type messageResponse struct {
	Message string `json:"message"`
}

func newMessageResponse(msg string) *messageResponse {
	return &messageResponse{
		Message: msg,
	}
}

func Response(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.Log(err.Error())
	}
}

func BadRequestErrResponse(w http.ResponseWriter, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	body := newMessageResponse(err.Error())
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.Log(err.Error())
	}
}
