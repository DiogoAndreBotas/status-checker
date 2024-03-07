package main

import (
	"log"
	"time"

	"github.com/jackc/pgx"
)

// Type definition

type statusCheck struct {
	StatusCode   int       `json:"status_code"`
	ResponseTime int       `json:"response_time"`
	Timestamp    time.Time `json:"timestamp"`
	Url          string    `json:"url"`
}

// Public methods

func GetStatusChecks(conn *pgx.Conn) []statusCheck {
	log.Default().Println("Getting status check from DB...")

	statusChecks := []statusCheck{}

	rows, err := conn.Query(`SELECT status_code, response_time, timestamp, url FROM StatusCheck`)
	if err != nil {
		log.Fatal("Failed to query rows from StatusCheck table: ", err)
	}

	for rows.Next() {
		var statusCheck statusCheck
		err = rows.Scan(&statusCheck.StatusCode, &statusCheck.ResponseTime, &statusCheck.Timestamp, &statusCheck.Url)
		if err != nil {
			log.Fatal("Failed to scan row from StatusCheck table: ", err)
		}

		statusChecks = append(statusChecks, statusCheck)
	}

	if rows.Err() != nil {
		log.Fatal("Failed to scan rows from StatusCheck table: ", rows.Err())
	}

	log.Default().Println("Got all status checks from DB!")

	return statusChecks
}

func GetLatestStatusChecks(conn *pgx.Conn) []statusCheck {
	log.Default().Println("Getting latest status checks from DB...")

	statusChecks := []statusCheck{}

	rows, err := conn.Query(`
		SELECT
			DISTINCT ON (url) url,
			status_code,
			response_time,
			timestamp
		FROM
			StatusCheck
		ORDER BY
			url,
			timestamp DESC;
	`)
	if err != nil {
		log.Fatal("Failed to query rows from StatusCheck table: ", err)
	}

	for rows.Next() {
		var statusCheck statusCheck
		err = rows.Scan(&statusCheck.Url, &statusCheck.StatusCode, &statusCheck.ResponseTime, &statusCheck.Timestamp)
		if err != nil {
			log.Fatal("Failed to scan row from StatusCheck table: ", err)
		}

		statusChecks = append(statusChecks, statusCheck)
	}

	if rows.Err() != nil {
		log.Fatal("Failed to scan rows from StatusCheck table: ", rows.Err())
	}

	log.Default().Println("Got latest status checks from DB!")

	return statusChecks
}

func AddUrl(conn *pgx.Conn, url string) string {
	_, err := conn.Exec(`INSERT INTO Url (url) VALUES ($1)`, url)
	if err != nil {
		log.Fatal("Failed to insert URL in Url table: ", err)
	}

	log.Default().Println("Added", url, "to Url table")

	// TODO return URL with ID (get query or check if exec returns anything)
	return url
}

func DeleteUrl(conn *pgx.Conn, id string) {
	_, err := conn.Exec(`DELETE FROM Url WHERE id=$1`, id)
	if err != nil {
		log.Fatal("Failed to delete URL from Url table: ", err)
	}

	log.Default().Println("Deleted URL with ID", id, "from Url table")
}
