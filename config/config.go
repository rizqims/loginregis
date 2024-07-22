package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
	Driver   string
}

type AppConfig struct {
	PortApp string
}

type Config struct {
	DBConfig
	AppConfig
}

func (c *Config) ReadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("env error: %v", err)
	}

	c.DBConfig = DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Dbname:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.AppConfig = AppConfig{
		PortApp: os.Getenv("PORTAPP"),
	}

	if c.User == "" || c.Password == "" || c.Host == "" || c.Port == "" || c.Dbname == "" || c.Driver == "" || c.PortApp == "" {
		return fmt.Errorf("ENV na kosong mang")
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.ReadConfig(); err != nil {
		return nil, fmt.Errorf("readconfig err: %v", err)
	}
	return config, nil
}
