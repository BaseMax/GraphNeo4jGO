package mux

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, code int, data any) error {
	enc := json.NewEncoder(w)
	w.WriteHeader(code)
	return enc.Encode(data)
}
