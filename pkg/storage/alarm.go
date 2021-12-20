package storage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Alarm struct {
	Alarm     bool      `json:"alarm"`
	Timestamp time.Time `json:"timestamp"`
}

type dbConfig struct {
	user    string
	pwd     string
	host    string
	port    string
	name    string
	connStr string
}

func newDbConfig() dbConfig {
	// Let's set some initial default variables
	c := dbConfig{
		user: os.Getenv("DB_USER"),
		pwd:  os.Getenv("DB_PWD"),
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
		name: os.Getenv("DB_NAME"),
	}
	c.connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.user, c.pwd, c.host, c.port, c.name)
	return c
}

func Connect() *pgx.Conn {
	cfg := newDbConfig()
	conn, err := pgx.Connect(context.Background(), cfg.connStr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Database connection failed.")
		os.Exit(1)
	}
	return conn
}

func GetLastAlarms(limit int) ([]Alarm, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	aa := []Alarm{}

	q := "select timestamp, alarm from alarms ORDER BY timestamp DESC LIMIT $1"
	rows, err := conn.Query(context.Background(), q, limit)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Read failed.")
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

func InsertAlarm(a Alarm) error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "INSERT INTO alarms VALUES ($1, $2)", a.Alarm, a.Timestamp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Insert failed.")
		return errors.New("Insert failed.")
	}
	logrus.Info("Insertion succesful.")
	return nil
}

func ClearAlarms() error {
	conn := Connect()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "DELETE FROM alarms")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Deletion failed.")
		return errors.New("Delete failed.")
	}
	logrus.Info("Clear succesful.")
	return nil
}
