package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"schilli.com/alarm-backend/pkg/storage"
)

func createAlarm(w http.ResponseWriter, r *http.Request, db Database) {
	var a storage.Alarm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		fmt.Println("Decode failed: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	a.Timestamp = time.Now()

	if err := storage.InsertAlarm(db.conn, a); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getAllAlarms(w http.ResponseWriter, r *http.Request, db Database) {
	aa, err := storage.GetLastAlarms(db.conn, 100)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aa)
}
