package main

import (
	"context"
	"golang/internal/handler"
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
	e.GET("/", handler.HelloWorldHandler)
	e.GET("/students/:id", handler.HandleGetStudent(conn))
	e.GET("/all_class_schedule", handler.HandleGetAllClassSchedules(conn))
	e.GET("/schedule/group/:id", handler.HandleGetScheduleByGroupID(conn))
	e.POST("/attendance/subject", handler.HandlePostSubjectAttendance(conn))
	e.GET("/attendanceBySubjectId/:id", handler.HandleGetAttendanceBySubjectID(conn))
	e.GET("/attendanceByStudentId/:id", handler.HandleGetAttendanceByStudentID(conn))

	e.Logger.Fatal(e.Start(":8080"))
}
