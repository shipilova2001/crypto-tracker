package shared

import (
	"encoding/json"
	"net/http"
)
func WriteJSON(w http.ResponseWriter, resp any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"code":500,"message":"internal server error"}`, http.StatusInternalServerError)
	}
}
