package main

import (
	"log"

	"github.com/joho/godotenv"
	"schilli.com/alarm-backend/pkg/rest"
	"schilli.com/alarm-backend/pkg/storage"
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// storage.GetLastAlarm()
	// storage.InsertAlarm(storage.Alarm{Timestamp: time.Now(), Alarm: false})
	// storage.GetLastAlarm()
	storage.ClearAlarms()
	storage.ClearActiveAlarms()
	// storage.GetLastAlarm()

	rest.StartServer("127.0.0.1", 8080)
}
