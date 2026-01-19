package repository

import (
	"context"
	"golang/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetStudentByID(connection *pgx.Conn, id int) (models.Student, error) {
	var student models.Student
	err := connection.QueryRow(context.Background(), "SELECT * FROM student WHERE student_id=$1", id).Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Email, &student.Gender, &student.BirthDate, &student.GroupID)

	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func GetGroupByID(connection *pgx.Conn, groupId int) (models.StudentGroup, error) {
	var group models.StudentGroup
	err := connection.QueryRow(context.Background(), "SELECT * FROM student_group WHERE group_id = $1", groupId).Scan(&group.GroupID, &group.FacultyID, &group.GroupName)

	if err != nil {
		return models.StudentGroup{}, err
	}

	return group, nil
}

func GetTimetables(connection *pgx.Conn) ([]models.Timetable, error) {
	rows, err := connection.Query(context.Background(), "SELECT * FROM timetable ORDER BY start_time ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// gathering data into a slice
	var timetables []models.Timetable
	for rows.Next() {
		var timetable models.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.CourseID, &timetable.Professor)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}

	return timetables, nil
}

func GetTimetableByGroupID(connection *pgx.Conn, groupId int) ([]models.Timetable, error) {
	var timetables []models.Timetable
	rows, err := connection.Query(context.Background(), "SELECT * FROM timetable WHERE group_id = $1 ORDER BY start_time ASC", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var timetable models.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.CourseID, &timetable.Professor)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}
	return timetables, nil
}

func RecordAttendance(connection *pgx.Conn, attendance models.AttendancePostRequest) error {
	_, err := connection.Exec(context.Background(), "INSERT INTO attendance (student_id, course_id, visited, visit_day) VALUES ($1, $2, $3, $4)", attendance.StudentID, attendance.CourseID, attendance.Visited, attendance.VisitDay)
	return err
}

func GetAttendanceBySubjectID(connection *pgx.Conn, courseId int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := connection.Query(context.Background(), "SELECT * FROM attendance WHERE course_id = $1", courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.AttendanceID, &attendance.StudentID, &attendance.CourseID, &attendance.Visited, &attendance.VisitDay)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func GetAttendanceByStudentID(connection *pgx.Conn, studentId int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := connection.Query(context.Background(), "SELECT * FROM attendance WHERE student_id = $1", studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.AttendanceID, &attendance.StudentID, &attendance.CourseID, &attendance.Visited, &attendance.VisitDay)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func CreateUser(connection *pgx.Conn, email string, hashedPassword string) error {
	_, err := connection.Exec(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
	return err
}

func GetUserByEmail(connection *pgx.Conn, email string) (string, error) {
	var user models.User
	rows, err := connection.Query(context.Background(), "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Email, &user.Password)
		if err != nil {
			return "", err
		}
	}
	return user.Password, nil
}
