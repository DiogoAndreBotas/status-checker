package main

func main() {
	conn := Connect()
	defer conn.Close()

	RunWebApi(conn)
}
