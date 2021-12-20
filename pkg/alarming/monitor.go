package alarming

import (
	"context"
	"fmt"
	"time"

	"schilli.com/alarm-backend/pkg/storage"
)

func MonitorAlarms() {
	conn := storage.Connect()
	defer conn.Close(context.Background())

	for {
		alarms, err := storage.GetAllActiveAlarms()
		if err != nil {
			fmt.Printf("Getting active alarms failed: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if len(alarms) == 0 {
			fmt.Printf("No active errors.\n")
			time.Sleep(10 * time.Second)
			continue
		}

		Call()
		storage.ClearActiveAlarms()
		time.Sleep(30 * time.Second)
	}

}
