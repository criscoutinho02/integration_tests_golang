package datasource

import (
	"database/sql"
	"fmt"
	"os"
)

func getPostgresConfig() string {
	host := os.Getenv("REC_POSTGRES_SERVICE_HOST")
	port := os.Getenv("REC_POSTGRES_SERVICE_PORT")
	user := os.Getenv("REC_POSTGRES_SERVICE_USER")
	password := os.Getenv("REC_POSTGRES_SERVICE_PASSWORD")
	database := os.Getenv("REC_POSTGRES_SERVICE_DATABASE")

	if password != "" {
		password = "password=" + password
	}

	return fmt.Sprintf("host=%s port=%s user=%s "+
		"%s dbname=%s sslmode=disable", host, port, user, password, database)
}

func NewPostgresConn() (*sql.DB, error) {
	conf := getPostgresConfig()
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return db, err
	}

	dbError := db.Ping()

	return db, dbError
}
