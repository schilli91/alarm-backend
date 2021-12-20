package storage

import (
	"context"
	"errors"
	"fmt"
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
			fmt.Println("Unable to read due to: ", err)
			return []Alarm{}, errors.New("Select failed.")
		}

		aa = append(aa, a)
	}

	// fmt.Printf("Fetched: %v\n", aa)
	return aa, nil
}

func InsertActiveAlarm(a Alarm) error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "INSERT INTO active_alarms VALUES ($1, $2)", a.Alarm, a.Timestamp)
	if err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return errors.New("Insert failed.")
	}
	fmt.Println("Insertion succesful")
	return nil
}

func ClearActiveAlarms() error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "DELETE FROM active_alarms")
	if err != nil {
		fmt.Println("Unable to delete due to: ", err)
		return errors.New("Delete failed.")
	}
	fmt.Println("Clear succesful")
	return nil
}
