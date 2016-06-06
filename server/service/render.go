package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	contentType     = "Content-Type"
	jsonContentType = "application/json; charset=utf-8"
)

var (
	ErrMissingQueryParam = errors.New("missing query param")
	ErrInvalidQueryParam = errors.New("invalid query param")
	ErrInvalidPathParam  = errors.New("invalid path param")

	statusMap = map[error]int{
		ErrMissingQueryParam: http.StatusBadRequest,
		ErrInvalidQueryParam: http.StatusBadRequest,
		ErrInvalidPathParam:  http.StatusNotFound,
	}
)

type InternalError interface {
	error
	InternalError() string
}

type errResponse struct {
	Error string `json:"error"`
}

func renderError(w http.ResponseWriter, r *http.Request, err error) {
	ie, ok := err.(InternalError)
	if ok && ie != nil {
		log.Printf("Error: %s, Request: %v", ie.InternalError(), r)
	}

	status := statusMap[err]
	if status == 0 {
		status = http.StatusInternalServerError
	}

	render(w, status, errResponse{err.Error()})
}

func render(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set(contentType, jsonContentType)

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
