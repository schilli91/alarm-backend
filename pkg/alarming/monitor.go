package alarming

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"schilli.com/alarm-backend/pkg/storage"
)

func MonitorAlarms() {
	conn := storage.Connect()
	defer conn.Close(context.Background())

	for {
		alarms, err := storage.GetAllActiveAlarms()
		if err != nil {
			logrus.Error("Getting active alarms failed: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if len(alarms) == 0 {
			logrus.Info("No active errors.\n")
			time.Sleep(10 * time.Second)
			continue
		}

		logrus.Info("Alarm is active. Calling!")
		Call()

		storage.ClearActiveAlarms()
		time.Sleep(30 * time.Second)
	}

}
