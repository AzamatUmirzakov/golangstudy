package models

type Student struct {
	StudentID int    `json:"student_id"`
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

type Course struct {
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
}

type Timetable struct {
	TimetableID int    `json:"timetable_id"`
	FacultyID   int    `json:"faculty_id"`
	GroupID     int    `json:"group_id"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Weekday     string `json:"weekday"`
	Location    string `json:"location"`
	CourseID    int    `json:"course_id"`
	Professor   string `json:"professor"`
}

type Attendance struct {
	AttendanceID int    `json:"attendance_id"`
	StudentID    int    `json:"student_id"`
	CourseID     int    `json:"course_id"`
	Visited      bool   `json:"visited"`
	VisitDay     string `json:"visit_day"`
}

type AttendancePostRequest struct {
	StudentID int    `json:"student_id"`
	CourseID  int    `json:"course_id"`
	Visited   bool   `json:"visited"`
	VisitDay  string `json:"visit_day"`
}

type GetStudentResponse struct {
	Student   Student `json:"student"`
	GroupName string  `json:"group_name"`
}
