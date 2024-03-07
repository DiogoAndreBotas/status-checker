package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx"
)

func RunCronJobScheduler(conn *pgx.Conn) {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal("CRON job scheduler failed to initialize with error: ", err)
	}

	defer s.Shutdown()

	c := make(chan string)

	_, err = s.NewJob(
		gocron.DurationJob(1*time.Minute),
		gocron.NewTask(runStatusChecker, conn, c),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		log.Fatal("CRON job failed to be defined with error: ", err)
	}

	s.Start()

	for m := range c {
		log.Default().Println(m)
	}
}

func runStatusChecker(conn *pgx.Conn, c chan string) {
	log.Default().Println("Running status checker job")

	urls := GetUrls(conn)

	// TODO: if there are no URLs, don't run the job

	for _, url := range urls {
		statusCheck(url, conn)
	}

	c <- "Status checker job complete!"
}

func statusCheck(url string, conn *pgx.Conn) {
	resp, _ := http.Get(url)

	// TODO: add response_time
	response_time := 1

	AddStatusCheck(conn, url, resp.StatusCode, response_time)

	fmt.Println(url + " has status code " + strconv.Itoa(resp.StatusCode))
}
