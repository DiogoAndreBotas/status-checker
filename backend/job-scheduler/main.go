package main

import service "database-service"

func main() {
	conn := service.Connect()
	defer conn.Close()

	LoadData(conn)
	RunCronJobScheduler(conn)
}
