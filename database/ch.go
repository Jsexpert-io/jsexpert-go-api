package database

import (
	"fmt"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"os"
)

func SetupDatabase() *ch.DB {
	host := os.Getenv("CLICKHOUSE_URL")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("CLICKHOUSE_NATIVE_PORT")
	if port == "" {
		port = "9000"
	}
	database := os.Getenv("CLICKHOUSE_DATABASE")
	if database == "" {
		database = "default"
	}
	db := ch.Connect(
		// clickhouse://<user>:<password>@<host>:<port>/<database>?sslmode=disable
		ch.WithDSN(fmt.Sprintf("clickhouse://%s:%s/%s?sslmode=disable", host, port, database)),
	)

	db.AddQueryHook(chdebug.NewQueryHook(
		chdebug.WithVerbose(true),
		chdebug.FromEnv("CHDEBUG"),
	))
	return db
}
