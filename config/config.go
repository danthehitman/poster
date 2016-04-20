package config

import (
	"fmt"
	"os"
	"strings"
)

var Database dbConfig
var ResourcePrefix string

type dbConfig struct {
	ConnectionString   string
}

func Init() {
	//user=poster password=admin dbname=poster sslmode=disable search_path=api
	var connectionString = "host=$POSTGRES_PORT_5432_TCP_ADDR user=postgres password=hpvse1 search_path=api dbname=poster sslmode=disable"
	// parse the env variable and set it properly
	if strings.Contains(connectionString, "$POSTGRES_PORT_5432_TCP_ADDR") {
		host, found := os.LookupEnv("POSTGRES_PORT_5432_TCP_ADDR")
		if (!found) {
			host = "localhost"
		}
		connectionString = strings.Replace(connectionString, "$POSTGRES_PORT_5432_TCP_ADDR", host, -1)
	}
	fmt.Println("[database] Using psql paramaters:", connectionString)

	// Construct a dbConfig object from envData
	Database = dbConfig{
		ConnectionString:   connectionString,
	}

	//ResourcePrefix = "poster/src/"
	ResourcePrefix = ""
}
