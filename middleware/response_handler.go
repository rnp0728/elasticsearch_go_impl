package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/elasticsearch_go_impl/src"
)

func ResponseHandler(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// setting the content-type header to "application/json"
		w.Header().Set("Content-Type", "application/json")
		err := handler(w, r)

		if err != nil {

			resp := src.NewInternalServerException(err)
			w.WriteHeader(http.StatusInternalServerError)

			jsonEncoder := json.NewEncoder(w)
			jsonEncoder.Encode(resp)
			return
		}
	}
}