package handler

import (
	"golang/internal/models"
	"golang/internal/repository"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "Hello World!"})
}

func HandleGetStudent(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		// student query
		student, err := repository.GetStudentByID(conn, id)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// group query
		group, err := repository.GetGroupByID(conn, student.GroupID)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// response
		var response = models.GetStudentResponse{
			Student:   student,
			GroupName: group.GroupName,
		}
		c.JSON(200, response)

		return nil
	}
}

func HandleGetAllClassSchedules(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		timetables, err := repository.GetTimetables(conn)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, timetables)
		return nil
	}
}

func HandleGetScheduleByGroupID(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid group id"})
		}

		timetables, err := repository.GetTimetableByGroupID(conn, groupId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, timetables)
		return nil
	}
}

func HandlePostSubjectAttendance(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var attendanceRequest models.AttendancePostRequest
		err := c.Bind(&attendanceRequest)

		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		err = repository.RecordAttendance(conn, attendanceRequest)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "attendance recorded successfully"})
	}
}

func HandleGetAttendanceByStudentID(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		attendances, err := repository.GetAttendanceByStudentID(conn, studentId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, attendances)
		return nil
	}
}

func HandleGetAttendanceBySubjectID(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		courseId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid course id"})
		}

		attendances, err := repository.GetAttendanceBySubjectID(conn, courseId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, attendances)
		return nil
	}
}
