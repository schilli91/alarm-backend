package main

import (
	"context"

	"schilli.com/alarm-backend/pkg/rest"
	"schilli.com/alarm-backend/pkg/storage"
)

func main() {
	conn := storage.Connect()
	defer conn.Close(context.Background())

	// storage.GetLastAlarm(conn)
	// storage.InsertAlarm(conn, storage.Alarm{Timestamp: time.Now(), Alarm: false})
	// storage.GetLastAlarm(conn)
	storage.ClearAlarms(conn)
	// storage.GetLastAlarm(conn)

	rest.StartServer("127.0.0.1", 8080, conn)
}
