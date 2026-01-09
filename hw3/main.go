package main

import (
	"context"
	"hw3/handlers"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	// start the server
	e := echo.New()
	e.GET("/students/:id", handlers.HandleGetStudent(conn))
	e.GET("/all_class_schedule", handlers.HandleGetAllClassSchedules(conn))
	e.GET("/schedule/group/:id", handlers.HandleGetScheduleByGroupID(conn))

	e.Logger.Fatal(e.Start(":8080"))
}
