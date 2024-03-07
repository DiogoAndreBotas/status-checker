package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

func Connect() *pgx.Conn {
	host, host_exists := os.LookupEnv("POSTGRES_HOST")
	if !host_exists {
		host = "localhost"
	}

	port, port_exists := os.LookupEnv("POSTGRES_PORT")
	if !port_exists {
		port = "5432"
	}

	port_int, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Conversion from string to int failed", err)
	}

	db, db_exists := os.LookupEnv("POSTGRES_DB")
	if !db_exists {
		db = "steam-status-checker"
	}

	user, user_exists := os.LookupEnv("POSTGRES_USER")
	if !user_exists {
		user = "dbotas"
	}

	pw, pw_exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !pw_exists {
		pw = "dbotaspass"
	}

	connConfig := pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port_int),
		Database: db,
		User:     user,
		Password: pw,
	}
	conn, err := pgx.Connect(connConfig)

	log.Default().Println("Connecting to DB instance with the following parameters", host, port, db, user)

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	log.Default().Println("Successfully connected to Postgres DB instance", host)

	return conn
}
