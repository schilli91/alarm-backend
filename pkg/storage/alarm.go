package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
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
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func GetLastAlarms(conn *pgx.Conn, limit int) ([]Alarm, error) {
	aa := []Alarm{}

	q := "select timestamp, alarm from alarms ORDER BY timestamp DESC LIMIT $1"
	rows, err := conn.Query(context.Background(), q, limit)
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

	fmt.Printf("Fetched: %v\n", aa)
	return aa, nil
}

func InsertAlarm(conn *pgx.Conn, a Alarm) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO alarms VALUES ($1, $2)", a.Alarm, a.Timestamp)
	if err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return errors.New("Insert failed.")
	}
	fmt.Println("Insertion succesful")
	return nil
}

func ClearAlarms(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM alarms")
	if err != nil {
		fmt.Println("Unable to delete due to: ", err)
		return errors.New("Delete failed.")
	}
	fmt.Println("Clear succesful")
	return nil
}
