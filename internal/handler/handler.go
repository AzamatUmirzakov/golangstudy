package handler

import (
	"golang/internal/models"
	"golang/internal/repository"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

func HandleUserRegister(conn *pgx.Conn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var registerRequest models.UserRegisterRequest
		err := c.Bind(&registerRequest)
		if err != nil {
			c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(500, map[string]string{"error": "failed to hash password"})
		}

		err = repository.CreateUser(conn, registerRequest.Email, string(hashedPassword))

		if err != nil {
			c.JSON(500, map[string]string{"error": err.Error()})
		}

		c.JSON(200, map[string]string{"message": "user registered successfully"})
		return nil
	}
}

func HandleUserLogin(conn *pgx.Conn, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest models.UserRegisterRequest
		err := c.Bind(&loginRequest)
		if err != nil {
			c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		hashedPassword, err := repository.GetUserByEmail(conn, loginRequest.Email)
		if err != nil {
			c.JSON(500, map[string]string{"error": err.Error()})
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginRequest.Password))
		if err != nil {
			c.JSON(401, map[string]string{"error": "invalid password"})
		}

		claims := models.UserClaims{
			Email: loginRequest.Email,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			c.JSON(500, map[string]string{"error": "failed to generate token"})
			return err
		}

		c.Response().Header().Set("Authorization", "Bearer "+tokenString)

		c.JSON(200, map[string]string{"message": "user logged in successfully"})
		return nil
	}
}
