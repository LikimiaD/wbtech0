package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) respondWithError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": status,
		"error":  err.Error(),
	})
}

func (s *Server) respondAny(w http.ResponseWriter, httpStatus int, obj any) {
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": httpStatus,
		"result": obj,
	})
}
