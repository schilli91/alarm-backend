package rest

import (
	"encoding/json"
	"net/http"

	"schilli.com/alarm-backend/pkg/storage"
)

func getAllActiveAlarms(w http.ResponseWriter, r *http.Request) {
	aa, err := storage.GetAllActiveAlarms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aa)
}
