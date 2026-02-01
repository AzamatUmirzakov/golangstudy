package main

import (
	"context"
	"golang/internal/handler"
	"golang/internal/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	// start the server
	e := echo.New()
	e.GET("/", handler.HelloWorldHandler)
	e.Use(middleware.RequestLogger())
	// cors config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.POST("api/auth/register", handler.HandleUserRegister(pool))
	e.POST("api/auth/login", handler.HandleUserLogin(pool, os.Getenv("JWT_SECRET")))

	api := e.Group("/")
	api.Use(JWTMiddleware)
	api.GET("student/:id", handler.HandleGetStudent(pool))
	api.GET("students", handler.HandleGetAllStudents(pool))
	api.GET("groups", handler.HandleGetGroups(pool))
	api.POST("groups", handler.HandlePostGroup(pool))
	api.PUT("groups/:id", handler.HandlePutGroup(pool))
	api.DELETE("groups/:id", handler.HandleDeleteGroup(pool))
	api.GET("faculties", handler.HandleGetFaculties(pool))
	api.POST("faculties", handler.HandlePostFaculty(pool))
	api.PUT("faculties/:id", handler.HandleUpdateFaculty(pool))
	api.DELETE("faculties/:id", handler.HandleDeleteFaculty(pool))
	api.GET("all_class_schedule", handler.HandleGetAllClassSchedules(pool))
	api.POST("student", handler.HandlePostStudent(pool))
	api.PUT("student/:id", handler.HandleUpdateStudent(pool))
	api.DELETE("student/:id", handler.HandleDeleteStudent(pool))
	api.GET("schedule/group/:id", handler.HandleGetScheduleByGroupID(pool))
	api.POST("attendance/subject", handler.HandlePostSubjectAttendance(pool))
	api.GET("attendanceBySubjectId/:id", handler.HandleGetAttendanceBySubjectID(pool))
	api.GET("attendanceByStudentId/:id", handler.HandleGetAttendanceByStudentID(pool))

	e.Logger.Fatal(e.Start(":8080"))
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(401, map[string]string{"error": "missing or invalid token"})
		}

		token := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		claims := &models.UserClaims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
		if err != nil || !parsedToken.Valid {
			return c.JSON(401, map[string]string{"error": "invalid token"})
		}

		c.Set("user_claims", claims)

		return next(c)
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, echo.NewHTTPError(401, "unexpected signing method")
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	return secret, nil
}
