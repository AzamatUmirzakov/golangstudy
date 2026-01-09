package handlers

import (
	"hw3/getters"
	"hw3/types"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func HandleGetStudent(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		// student query
		student, err := getters.GetStudentByID(conn, id)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// group query
		group, err := getters.GetGroupByID(conn, student.GroupID)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// response
		var response = types.GetStudentResponse{
			Student:   student,
			GroupName: group.GroupName,
		}
		c.JSON(200, response)

		return nil
	}
}

func HandleGetAllClassSchedules(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		timetables, err := getters.GetTimetables(conn)
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

		timetables, err := getters.GetTimetableByGroupID(conn, groupId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, timetables)
		return nil
	}
}
