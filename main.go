package main

import (
	"context"
	"docker-study/db"
	"docker-study/employees"
	"docker-study/networking"
	"fmt"
)

func main() {
	conn, err := db.ConnectToDB(context.Background())
	if err != nil {
		panic(err)
	}
	handlers := networking.NewEmployeesHandlers(context.Background(), conn)
	server := networking.NewEmployeesServer(handlers)

	if err := employees.CreateTable(context.Background(), conn); err != nil {
		panic(err)
	}

	if err := server.StartServer(":5555"); err != nil {
		fmt.Println(err)
	}
}
