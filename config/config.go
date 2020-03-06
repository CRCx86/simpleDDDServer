package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {

	Hostname := "localhost"
	HostPort := "5432"
	Username := "postgres"
	Password := "postgres"
	DatabaseName := "postgres"

	DatabasePath := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Hostname, HostPort, Username, Password, DatabaseName)

	return &Config{
		Enabled:      true,
		DatabasePath: DatabasePath,
		Port:         ":8081",
	}
}

func (c *Config) Connect(config *Config) (*sql.DB, error) {
	return sql.Open("postgres", config.DatabasePath)
}
