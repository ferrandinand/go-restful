package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type appErr struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func JsonError(w http.ResponseWriter, handlerErr error, code int) {
	log.Printf("Error: %s\n", handlerErr)

	w.Header().Set("Content-Type", "application/json; chaset=utf-8")
	w.WriteHeader(code)
	if res, err := json.Marshal(&appErr{handlerErr.Error(), code}); err == nil {
		w.Write(res)
	}
}

func JsonOk(w http.ResponseWriter, res []byte, code int) {

	w.Header().Set("Content-Type", "application/json; chaset=utf-8")
	w.WriteHeader(code)
	w.Write(res)
}

func JsonStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
