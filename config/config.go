package config

import (
	"fmt"
	"log"
	"os"
)

type dataBaseConfig struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

var c dataBaseConfig

// "user:pass@tcp(ip:port)/name?charset=utf8&parseTime=True&loc=Local"
const connectionStringTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

func init() {
	c = dataBaseConfig {
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		IP: os.Getenv("DB_IP"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	err := isValidConfig(c)

	if err != nil {
		log.Fatal(err)
	}
}

func isValidConfig(c dataBaseConfig) error {
	if c.User == "" {
		return fmt.Errorf("db user isn't set")
	}
	if c.Pass == "" {
		return fmt.Errorf("db pass isn't set")
	}
	if c.IP == "" {
		return fmt.Errorf("db ip isn't set")
	}
	if c.Port == "" {
		return fmt.Errorf("db port isn't set")
	}
	if c.Name == "" {
		return fmt.Errorf("db name isn't set")
	}
	return nil
}

func GetConnectionString() string {
	return fmt.Sprintf(connectionStringTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}