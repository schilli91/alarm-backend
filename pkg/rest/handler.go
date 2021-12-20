package rest

import (
	"fmt"
	"net/http"
)

func StartServer(host string, port int) {
	http.HandleFunc("/", hello)
	http.HandleFunc("/alarms", alarms)
	http.HandleFunc("/active-alarms", activeAlarms)

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

func alarms(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createAlarm(w, r)
		return
	}

	if r.Method == "GET" {
		getAllAlarms(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("<h1>Method Not Allowed</h1>"))
	return
}

func activeAlarms(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {
	// 	createAlarm(w, r)
	// 	return
	// }

	if r.Method == "GET" {
		getAllActiveAlarms(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("<h1>Method Not Allowed</h1>"))
	return
}
