package main

import (
	"log"
	"time"

	"github.com/jackc/pgx"
)

// Public methods

func LoadData(conn *pgx.Conn) {
	createUrlTable(conn)
	createStatusCheckTable(conn)
}

func GetUrls(conn *pgx.Conn) []string {
	urls := []string{}

	rows, err := conn.Query(`SELECT url FROM Url`)
	if err != nil {
		log.Fatal("Failed to query rows from Url table: ", err)
	}

	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		if err != nil {
			log.Fatal("Failed to scan row from Url table: ", err)
		}

		log.Default().Println("Got", url, "from Url table")

		urls = append(urls, url)
	}

	if rows.Err() != nil {
		log.Fatal("Failed to scan row from Url table: ", rows.Err())
	}

	return urls
}

func AddStatusCheck(conn *pgx.Conn, url string, statusCode int, responseTime int) {
	ts := time.Now()
	_, err := conn.Exec(`
		INSERT INTO
			StatusCheck (url, status_code, response_time, timestamp)
		VALUES
			($1, $2, $3, $4)
	`, url, statusCode, responseTime, ts)
	if err != nil {
		log.Fatal("Failed to insert URLs in StatusCheck table: ", err)
	}

	log.Default().Println("Added StatusCheck with following details", url, statusCode, responseTime, ts)
}

// Private methods

// TODO: find a way of loading this data in another way
// problems could occur if the web_api boots up before the data loader
// maybe separate to a diff service?
func createUrlTable(conn *pgx.Conn) {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS Url (
			id SERIAL PRIMARY KEY,
			url TEXT UNIQUE
		)
	`)
	if err != nil {
		log.Fatal("Failed to create Url table: ", err)
	}

	log.Default().Println("Created Url table")
}

func createStatusCheckTable(conn *pgx.Conn) {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS StatusCheck (
			id SERIAL PRIMARY KEY,
			status_code SMALLINT NOT NULL,
			response_time SMALLINT NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			url TEXT,
			CONSTRAINT fk_url FOREIGN KEY(url) REFERENCES Url(url)
		)
	`)
	if err != nil {
		log.Fatal("Failed to create StatusCheck table: ", err)
	}

	log.Default().Println("Created StatusCheck table")
}
