package config

import (
	"fmt"
	"os"
	"strings"
)

var Database dbConfig

type config struct {
	Database dbConfig
}

type dbConfig struct {
	ConnectionString   string
}

func Init() {
	ParseDatabaseConfig("development")
}

func ParseDatabaseConfig(env string) {
	//user=poster password=admin dbname=poster sslmode=disable search_path=api
	var connectionString = "host=$POSTGRES_PORT_5432_TCP_ADDR user=poster search_path=api dbname=poster sslmode=disable"
	// parse the env variable and set it properly
	if strings.Contains(connectionString, "$POSTGRES_PORT_5432_TCP_ADDR") {
		connectionString = strings.Replace(connectionString, "$POSTGRES_PORT_5432_TCP_ADDR", os.Getenv("POSTGRES_PORT_5432_TCP_ADDR"), -1)
	}
	fmt.Println("[database] Using psql paramaters:", connectionString)

	// Construct a dbConfig object from envData
	Database = dbConfig{
		ConnectionString:   connectionString,
	}
}
