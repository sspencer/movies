package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func replyJSON(w http.ResponseWriter, r *http.Request, status int, resp interface{}) {

	body, err := json.Marshal(resp)

	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("%s %s HTTP %d", r.Method, r.URL, status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(body))
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error %d %q", r.URL, r.Method, status, message)
	http.Error(w, message, status)
}
