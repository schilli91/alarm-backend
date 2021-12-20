package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewFromJSON(name string) Conf {
	f, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Reading file failed: %v\n", err)
		os.Exit(1)
	}

	conf := Conf{}
	json.Unmarshal(f, &conf)
	fmt.Printf("Config set to: %v\n", conf)
	return conf
}
