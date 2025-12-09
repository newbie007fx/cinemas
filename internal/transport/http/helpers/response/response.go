package response

import (
	"encoding/json"
	"net/http"

	"github.com/newbie007fx/cinemas/internal/errors"
)

type Response[T any] struct {
	IsSuccess bool              `json:"is_success"`
	Data      T                 `json:"data,omitempty"`
	Error     *errors.BaseError `json:"error,omitempty"`
}

func (resp Response[T]) Send(wr http.ResponseWriter) {
	statusCode := 200
	if resp.Error != nil {
		statusCode = resp.Error.GetStatusCode()
	}

	wr.Header().Set("Content-type", "application/json; charset=utf-8")
	wr.WriteHeader(statusCode)
	payload, _ := json.Marshal(resp)
	wr.Write(payload)
}

func WriterResponseError(wr http.ResponseWriter, err *errors.BaseError) {
	resp := Response[map[string]string]{
		Error: err,
	}

	resp.Send(wr)
}
