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

type StudentPostRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	GroupID   int    `json:"group_id"`
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

type Subject struct {
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	FacultyID   int    `json:"faculty_id"`
	ProfessorID int    `json:"professor_id"`
}

type Timetable struct {
	TimetableID int       `json:"timetable_id"`
	FacultyID   int       `json:"faculty_id"`
	GroupID     int       `json:"group_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Weekday     string    `json:"weekday"`
	Location    string    `json:"location"`
	SubjectID   int       `json:"subject_id"`
}

type Attendance struct {
	AttendanceID int       `json:"attendance_id"`
	StudentID    int       `json:"student_id"`
	TimetableID  int       `json:"timetable_id"`
	SubjectID    int       `json:"subject_id"`
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

type Professor struct {
	ProfessorID int    `json:"professor_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	FacultyID   int    `json:"faculty_id"`
}
