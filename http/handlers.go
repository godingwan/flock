package http

import (
	"encoding/json"
	"net/http"

	"github.com/jsungholee/flock"
)

// HandleSearch handles request for searching tweets
func HandleSearch(service flock.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		ht, ok := params["hashtag"]
		if !ok {
			badRequest(w, "No hashtag to search given")
		}

		tweets, err := service.GetTweets(r.Context(), ht[0])
		if err != nil {
			serverError(w, "Failed to get tweets")
		}

		data, err := json.Marshal(tweets)
		if err != nil {
			serverError(w, "Failed to marshal data")
		}

		statusOK(w, data)
	})
}

//======================== Responses ========================
// 2xx responses
func statusOK(w http.ResponseWriter, data []byte) (int, error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w.Write(data)
}

// 4xx responses
func badRequest(w http.ResponseWriter, err string) (int, error) {
	w.WriteHeader(http.StatusBadRequest)
	return w.Write([]byte(err))
}

// 5xx responses
func serverError(w http.ResponseWriter, err string) (int, error) {
	w.WriteHeader(http.StatusBadRequest)
	return w.Write([]byte(err))
}
