package mux

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, code int, data any) {
	enc := json.NewEncoder(w)
	w.WriteHeader(code)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
}
