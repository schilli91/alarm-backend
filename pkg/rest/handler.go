package rest

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func StartServer(host string, port int, conn *pgx.Conn) {
	db := Database{conn: conn}

	http.HandleFunc("/", hello)
	http.HandleFunc("/alarms", db.alarms)

	//conf := NewFromJSON("config.json")

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n", r.Method)
	fmt.Fprintf(w, "Hello %s", "world")
}

func (db Database) alarms(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createAlarm(w, r, db)
		return
	}

	if r.Method == "GET" {
		getAllAlarms(w, r, db)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("<h1>Method Not Allowed</h1>"))
	return
}
