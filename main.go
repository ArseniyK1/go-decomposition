package main

import (
	"context"
	"decomposition/internal/app/middleware"
	"decomposition/internal/app/routers"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {

	router := http.NewServeMux()
	routers.Router(router)

	conn, err := pgx.Connect(context.Background(), "postgres://posgtres:root@localhost:5432/dbname")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to database!")

	fmt.Println("Server is running on http://localhost:8080")

	error := http.ListenAndServe(":8080", middleware.Middleware(router))

	if error != nil {
		fmt.Println("Server is not running", err)
	}
}
