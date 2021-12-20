package storage

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func GetAllActiveAlarms() ([]Alarm, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	aa := []Alarm{}

	q := "select timestamp, alarm from active_alarms ORDER BY timestamp DESC"
	rows, err := conn.Query(context.Background(), q)
	if err != nil {
		return aa, err
	}

	for rows.Next() {
		var a Alarm
		err := rows.Scan(&a.Timestamp, &a.Alarm)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Error": err,
			}).Error("Read failed.")
			return []Alarm{}, errors.New("Select failed.")
		}

		aa = append(aa, a)
	}

	return aa, nil
}

func InsertActiveAlarm(a Alarm) error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "INSERT INTO active_alarms VALUES ($1, $2)", a.Alarm, a.Timestamp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Insert failed.")
		return errors.New("Insert failed.")
	}
	logrus.WithFields(logrus.Fields{
		"Alarm": a,
	}).Info("Insertion succesful.")
	return nil
}

func ClearActiveAlarms() error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "DELETE FROM active_alarms")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Deletion failed.")
		return errors.New("Delete failed.")
	}
	logrus.Info("Clear succesful.")
	return nil
}
