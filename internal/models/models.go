package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Student struct {
	StudentID int       `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	GroupID   int       `json:"group_id"`
}

type StudentGroup struct {
	GroupID   int    `json:"group_id"`
	FacultyID int    `json:"faculty_id"`
	GroupName string `json:"group_name"`
}

type Faculty struct {
	FacultyID   int    `json:"faculty_id"`
	FacultyName string `json:"faculty_name"`
}

type Course struct {
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	FacultyID  int    `json:"faculty_id"`
}

type Timetable struct {
	TimetableID int       `json:"timetable_id"`
	FacultyID   int       `json:"faculty_id"`
	GroupID     int       `json:"group_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Weekday     string    `json:"weekday"`
	Location    string    `json:"location"`
	CourseID    int       `json:"course_id"`
}

type Attendance struct {
	AttendanceID int       `json:"attendance_id"`
	StudentID    int       `json:"student_id"`
	TimetableID  int       `json:"timetable_id"`
	CourseID     int       `json:"course_id"`
	Visited      bool      `json:"visited"`
	VisitDay     time.Time `json:"visit_day"`
}

type AttendancePostRequest struct {
	StudentID   int    `json:"student_id"`
	TimetableID int    `json:"timetable_id"`
	Visited     bool   `json:"visited"`
	VisitDay    string `json:"visit_day"`
}

type GetStudentResponse struct {
	Student   Student `json:"student"`
	GroupName string  `json:"group_name"`
}

type User struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
