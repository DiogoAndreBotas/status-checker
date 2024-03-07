package main

func main() {
	conn := Connect()
	defer conn.Close()

	LoadData(conn)
	RunCronJobScheduler(conn)
}
