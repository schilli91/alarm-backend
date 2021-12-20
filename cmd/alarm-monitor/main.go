package main

import (
	"log"

	"github.com/joho/godotenv"
	"schilli.com/alarm-backend/pkg/alarming"
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alarming.MonitorAlarms()
}
