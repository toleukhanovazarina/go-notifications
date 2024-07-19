package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func main() {
	connString := "postgresql://zarinatoleukhanova:123456@localhost:5432/go-notification"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to PostgreSQL!")
}
