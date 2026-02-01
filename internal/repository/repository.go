package repository

import (
	"context"
	"golang/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllStudents(pool *pgxpool.Pool) ([]models.Student, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM student ORDER BY student_id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Email, &student.Gender, &student.BirthDate, &student.GroupID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func GetAllGroups(pool *pgxpool.Pool) ([]models.StudentGroup, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM student_group")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.StudentGroup
	for rows.Next() {
		var group models.StudentGroup
		err := rows.Scan(&group.GroupID, &group.FacultyID, &group.GroupName)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func GetAllFaculties(pool *pgxpool.Pool) ([]models.Faculty, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM faculty")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculties []models.Faculty
	for rows.Next() {
		var faculty models.Faculty
		err := rows.Scan(&faculty.FacultyID, &faculty.FacultyName)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return faculties, nil
}

func GetStudentByID(pool *pgxpool.Pool, id int) (models.Student, error) {
	var student models.Student
	err := pool.QueryRow(context.Background(), "SELECT * FROM student WHERE student_id=$1", id).Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Email, &student.Gender, &student.BirthDate, &student.GroupID)

	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func GetGroupByID(pool *pgxpool.Pool, groupId int) (models.StudentGroup, error) {
	var group models.StudentGroup
	err := pool.QueryRow(context.Background(), "SELECT * FROM student_group WHERE group_id = $1", groupId).Scan(&group.GroupID, &group.FacultyID, &group.GroupName)

	if err != nil {
		return models.StudentGroup{}, err
	}

	return group, nil
}

func GetTimetableByTimetableID(pool *pgxpool.Pool, timetableId int) (models.Timetable, error) {
	var timetable models.Timetable
	err := pool.QueryRow(context.Background(), "SELECT * FROM timetable WHERE timetable_id = $1", timetableId).Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.CourseID)
	if err != nil {
		return models.Timetable{}, err
	}

	return timetable, nil
}

func GetTimetables(pool *pgxpool.Pool) ([]models.Timetable, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM timetable ORDER BY start_time ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// gathering data into a slice
	var timetables []models.Timetable
	for rows.Next() {
		var timetable models.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.CourseID)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return timetables, nil
}

func GetTimetableByGroupID(pool *pgxpool.Pool, groupId int) ([]models.Timetable, error) {
	var timetables []models.Timetable
	rows, err := pool.Query(context.Background(), "SELECT * FROM timetable WHERE group_id = $1 ORDER BY start_time ASC", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var timetable models.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.CourseID)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return timetables, nil
}

func RecordAttendance(pool *pgxpool.Pool, attendance models.Attendance) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO attendance (student_id, timetable_id, course_id, visited, visit_day) VALUES ($1, $2, $3, $4, $5)", attendance.StudentID, attendance.TimetableID, attendance.CourseID, attendance.Visited, attendance.VisitDay)
	return err
}

func GetAttendanceBySubjectID(pool *pgxpool.Pool, courseId int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := pool.Query(context.Background(), "SELECT * FROM attendance WHERE course_id = $1", courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.AttendanceID, &attendance.Visited, &attendance.VisitDay, &attendance.StudentID, &attendance.CourseID, &attendance.TimetableID)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func GetAttendanceByStudentID(pool *pgxpool.Pool, studentId int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := pool.Query(context.Background(), "SELECT * FROM attendance WHERE student_id = $1", studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.AttendanceID, &attendance.Visited, &attendance.VisitDay, &attendance.StudentID, &attendance.CourseID, &attendance.TimetableID)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func CreateUser(pool *pgxpool.Pool, email string, hashedPassword string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
	return err
}

func GetUserByEmail(pool *pgxpool.Pool, email string) (string, error) {
	var user models.User
	err := pool.QueryRow(context.Background(), "SELECT user_id, email, password FROM users WHERE email = $1", email).Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func CreateStudent(pool *pgxpool.Pool, student models.StudentPostRequest) (int, error) {
	var studentID int
	err := pool.QueryRow(context.Background(), "INSERT INTO student (first_name, last_name, email, gender, birth_date, group_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING student_id", student.FirstName, student.LastName, student.Email, student.Gender, student.BirthDate, student.GroupID).Scan(&studentID)
	if err != nil {
		return 0, err
	}
	return studentID, nil
}

func UpdateStudent(pool *pgxpool.Pool, id int, student models.StudentPostRequest) error {
	_, err := pool.Exec(context.Background(), "UPDATE student SET first_name=$1, last_name=$2, email=$3, gender=$4, birth_date=$5, group_id=$6 WHERE student_id=$7", student.FirstName, student.LastName, student.Email, student.Gender, student.BirthDate, student.GroupID, id)
	return err
}

func DeleteStudent(pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM student WHERE student_id=$1", id)
	return err
}
