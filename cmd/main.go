package main

import (
	"context"
	"golang/internal/handler"
	"golang/internal/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
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

	e.POST("api/auth/register", handler.HandleUserRegister(conn))
	e.POST("api/auth/login", handler.HandleUserLogin(conn, os.Getenv("JWT_SECRET")))

	api := e.Group("/")
	api.Use(JWTMiddleware)
	api.GET("students/:id", handler.HandleGetStudent(conn))
	api.GET("all_class_schedule", handler.HandleGetAllClassSchedules(conn))
	api.GET("schedule/group/:id", handler.HandleGetScheduleByGroupID(conn))
	api.POST("attendance/subject", handler.HandlePostSubjectAttendance(conn))
	api.GET("attendanceBySubjectId/:id", handler.HandleGetAttendanceBySubjectID(conn))
	api.GET("attendanceByStudentId/:id", handler.HandleGetAttendanceByStudentID(conn))

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
