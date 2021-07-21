package config

import (
	"encoding/json"
	"github.com/kislball/dva/pkg/database"
	"os"
)

// a common config is used for all services because i want so

type Config struct {
	Discord struct{
		Token string `json:"token"`
		Secret string `json:"secret"`
		ClientID string `json:"client_id"`
	} `json:"discord"`
	NATS struct{
		URI string `json:"uri"`
	} `json:"nats"`
	Database struct{
		DSN string `json:"dsn"`
		Driver database.Driver `json:"driver"`
	} `json:"database"`
	Gateway struct{
		APIPort string `json:"port"`
	} `json:"gateway"`
}

func (c *Config) Load(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	d := json.NewDecoder(f)
	err = d.Decode(c)
	return
}