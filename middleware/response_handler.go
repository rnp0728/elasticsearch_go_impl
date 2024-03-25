package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/elasticsearch_go_impl/src"
)

// ResponseHandler is a middleware function that handles responses
// by setting the content-type header to "application/json" and
// encoding the response as JSON.
func ResponseHandler(handler func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // setting the content-type header to "application/json"
        w.Header().Set("Content-Type", "application/json")
        err := handler(w, r)

        if err!= nil {
            // creating a new internal server exception
            resp := src.NewInternalServerException(err)
            // setting the status code to 500 internal server error
            w.WriteHeader(http.StatusInternalServerError)
            // encoding the response as JSON
            jsonEncoder := json.NewEncoder(w)
            jsonEncoder.Encode(resp)
            return
        }
    }
}