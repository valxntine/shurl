package rest

import (
	"errors"
	"io"
	"net/http"
)

func parseRequest(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("error reading request body")
	}
	return body, nil
}
