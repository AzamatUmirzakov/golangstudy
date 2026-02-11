package handler

import (
	"golang/internal/models"
	"golang/internal/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HelloWorldHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "Hello World!"})
}

func HandleGetStudent(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		// student query
		student, err := repository.GetStudentByID(pool, id)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// group query
		group, err := repository.GetGroupByID(pool, student.GroupID)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		// response
		var response = models.GetStudentResponse{
			Student:   student,
			GroupName: group.GroupName,
		}
		return c.JSON(200, response)
	}
}

func HandleGetAllStudents(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		students, err := repository.GetAllStudents(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, students)
	}
}

func HandleGetGroups(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		groups, err := repository.GetAllGroups(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, groups)
	}
}

func HandlePutGroup(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid group id"})
		}

		var group models.StudentGroup
		err = c.Bind(&group)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		err = repository.UpdateGroup(pool, id, group)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "group updated successfully"})
	}
}

func HandleDeleteGroup(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid group id"})
		}

		err = repository.DeleteGroup(pool, id)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "group deleted successfully"})
	}
}

func HandlePostGroup(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var group models.StudentGroup
		err := c.Bind(&group)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		id, err := repository.CreateGroup(pool, group)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "group created successfully", "group_id": strconv.Itoa(id)})
	}
}

func HandleGetFaculties(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		faculties, err := repository.GetAllFaculties(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, faculties)
	}
}

func HandlePostFaculty(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var faculty models.Faculty
		err := c.Bind(&faculty)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		id, err := repository.CreateFaculty(pool, faculty)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "faculty created successfully", "faculty_id": strconv.Itoa(id)})
	}
}

func HandleUpdateFaculty(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid faculty id"})
		}

		var faculty models.Faculty
		err = c.Bind(&faculty)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		err = repository.UpdateFaculty(pool, id, faculty)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "faculty updated successfully"})
	}
}

func HandleDeleteFaculty(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid faculty id"})
		}

		err = repository.DeleteFaculty(pool, id)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "faculty deleted successfully"})
	}
}

func HandleGetAllClassSchedules(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		timetables, err := repository.GetTimetables(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, timetables)
	}
}

func HandleGetScheduleByGroupID(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid group id"})
		}

		timetables, err := repository.GetTimetableByGroupID(pool, groupId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, timetables)
	}
}

func HandlePostStudent(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var student models.StudentPostRequest
		err := c.Bind(&student)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		id, err := repository.CreateStudent(pool, student)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "student created successfully ", "student_id": strconv.Itoa(id)})
	}
}

func HandleUpdateStudent(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		var student models.StudentPostRequest
		err = c.Bind(&student)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		err = repository.UpdateStudent(pool, studentId, student)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "student updated successfully"})
	}
}

func HandleDeleteStudent(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		err = repository.DeleteStudent(pool, studentId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "student deleted successfully"})
	}
}

func HandlePostSubjectAttendance(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var attendanceRequest models.AttendancePostRequest
		err := c.Bind(&attendanceRequest)

		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		var attendance models.Attendance
		attendance.StudentID = attendanceRequest.StudentID
		attendance.TimetableID = attendanceRequest.TimetableID
		attendance.Visited = attendanceRequest.Visited
		attendance.VisitDay, err = time.Parse("2006-01-02", attendanceRequest.VisitDay)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid date format"})
		}

		timetable, err := repository.GetTimetableByTimetableID(pool, attendance.TimetableID)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		attendance.SubjectID = timetable.SubjectID

		err = repository.RecordAttendance(pool, attendance)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "attendance recorded successfully"})
	}
}

func HandleGetAttendanceByStudentID(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid student id"})
		}

		attendances, err := repository.GetAttendanceByStudentID(pool, studentId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, attendances)
	}
}

func HandleGetAttendanceBySubjectID(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		courseId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid course id"})
		}

		attendances, err := repository.GetAttendanceBySubjectID(pool, courseId)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, attendances)
	}
}

func HandleUserRegister(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var registerRequest models.UserRegisterRequest
		err := c.Bind(&registerRequest)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

		if err != nil {
			return c.JSON(500, map[string]string{"error": "failed to hash password"})
		}

		err = repository.CreateUser(pool, registerRequest.Email, string(hashedPassword))

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"message": "user registered successfully"})
	}
}

func HandleUserLogin(pool *pgxpool.Pool, jwtSecret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest models.UserRegisterRequest
		err := c.Bind(&loginRequest)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request body"})
		}

		hashedPassword, err := repository.GetUserByEmail(pool, loginRequest.Email)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginRequest.Password))
		if err != nil {
			return c.JSON(401, map[string]string{"error": "invalid password"})
		}

		claims := models.UserClaims{
			Email: loginRequest.Email,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			return c.JSON(500, map[string]string{"error": "failed to generate token"})
		}

		return c.JSON(200, map[string]string{"message": "user logged in successfully", "token": tokenString})
	}
}

func HandleGetSubjects(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		subjects, err := repository.GetAllSubjects(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, subjects)
	}
}

func HandleGetProfessors(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		professors, err := repository.GetAllProfessors(pool)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, professors)
	}
}
