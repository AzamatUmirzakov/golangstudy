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
	rows, err := pool.Query(context.Background(), "SELECT * FROM student_group ORDER BY group_id ASC")
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

func CreateGroup(pool *pgxpool.Pool, group models.StudentGroup) (int, error) {
	var groupID int
	err := pool.QueryRow(context.Background(), "INSERT INTO student_group (faculty_id, group_name) VALUES ($1, $2) RETURNING group_id", group.FacultyID, group.GroupName).Scan(&groupID)
	if err != nil {
		return 0, err
	}
	return groupID, nil
}

func UpdateGroup(pool *pgxpool.Pool, id int, group models.StudentGroup) error {
	_, err := pool.Exec(context.Background(), "UPDATE student_group SET faculty_id=$1, group_name=$2 WHERE group_id=$3", group.FacultyID, group.GroupName, id)
	return err
}

func DeleteGroup(pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM student_group WHERE group_id=$1", id)
	return err
}

func GetAllFaculties(pool *pgxpool.Pool) ([]models.Faculty, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM faculty ORDER BY faculty_id ASC")
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

func CreateFaculty(pool *pgxpool.Pool, faculty models.Faculty) (int, error) {
	var facultyID int
	err := pool.QueryRow(context.Background(), "INSERT INTO faculty (faculty_name) VALUES ($1) RETURNING faculty_id", faculty.FacultyName).Scan(&facultyID)
	if err != nil {
		return 0, err
	}
	return facultyID, nil
}

func UpdateFaculty(pool *pgxpool.Pool, id int, faculty models.Faculty) error {
	_, err := pool.Exec(context.Background(), "UPDATE faculty SET faculty_name=$1 WHERE faculty_id=$2", faculty.FacultyName, id)
	return err
}

func DeleteFaculty(pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM faculty WHERE faculty_id=$1", id)
	return err
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
	err := pool.QueryRow(context.Background(), "SELECT * FROM timetable WHERE timetable_id = $1", timetableId).Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.SubjectID)
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
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.SubjectID)
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
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.SubjectID)
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
	_, err := pool.Exec(context.Background(), "INSERT INTO attendance (student_id, timetable_id, subject_id, visited, visit_day) VALUES ($1, $2, $3, $4, $5)", attendance.StudentID, attendance.TimetableID, attendance.SubjectID, attendance.Visited, attendance.VisitDay)
	return err
}

func GetAttendanceBySubjectID(pool *pgxpool.Pool, subjectId int) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := pool.Query(context.Background(), "SELECT * FROM attendance WHERE subject_id = $1", subjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.AttendanceID, &attendance.Visited, &attendance.VisitDay, &attendance.StudentID, &attendance.SubjectID, &attendance.TimetableID)
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
		err := rows.Scan(&attendance.AttendanceID, &attendance.Visited, &attendance.VisitDay, &attendance.StudentID, &attendance.SubjectID, &attendance.TimetableID)
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

func GetAllSubjects(pool *pgxpool.Pool) ([]models.Subject, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM subject ORDER BY subject_id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []models.Subject
	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(&subject.SubjectID, &subject.SubjectName, &subject.FacultyID, &subject.ProfessorID)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}

func GetSubjectByID(pool *pgxpool.Pool, id int) (models.Subject, error) {
	var subject models.Subject
	err := pool.QueryRow(context.Background(), "SELECT * FROM subject WHERE subject_id=$1", id).Scan(&subject.SubjectID, &subject.SubjectName, &subject.FacultyID, &subject.ProfessorID)

	if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func CreateSubject(pool *pgxpool.Pool, subject models.Subject) (int, error) {
	var subjectID int
	err := pool.QueryRow(context.Background(), "INSERT INTO subject (subject_name, faculty_id, professor_id) VALUES ($1, $2, $3) RETURNING subject_id", subject.SubjectName, subject.FacultyID, subject.ProfessorID).Scan(&subjectID)
	if err != nil {
		return 0, err
	}
	return subjectID, nil
}

func UpdateSubject(pool *pgxpool.Pool, id int, subject models.Subject) error {
	_, err := pool.Exec(context.Background(), "UPDATE subject SET subject_name=$1, faculty_id=$2, professor_id=$3 WHERE subject_id=$4", subject.SubjectName, subject.FacultyID, subject.ProfessorID, id)
	return err
}

func DeleteSubject(pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM subject WHERE subject_id=$1", id)
	return err
}

func GetAllProfessors(pool *pgxpool.Pool) ([]models.Professor, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM professor ORDER BY professor_id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var professors []models.Professor
	for rows.Next() {
		var professor models.Professor
		err := rows.Scan(&professor.ProfessorID, &professor.FirstName, &professor.LastName, &professor.Email, &professor.FacultyID)
		if err != nil {
			return nil, err
		}
		professors = append(professors, professor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return professors, nil
}

func GetProfessorByID(pool *pgxpool.Pool, id int) (models.Professor, error) {
	var professor models.Professor
	err := pool.QueryRow(context.Background(), "SELECT * FROM professor WHERE professor_id=$1", id).Scan(&professor.ProfessorID, &professor.FirstName, &professor.LastName, &professor.Email, &professor.FacultyID)

	if err != nil {
		return models.Professor{}, err
	}

	return professor, nil
}

func CreateProfessor(pool *pgxpool.Pool, professor models.Professor) (int, error) {
	var professorID int
	err := pool.QueryRow(context.Background(), "INSERT INTO professor (first_name, last_name, email, faculty_id) VALUES ($1, $2, $3, $4) RETURNING professor_id", professor.FirstName, professor.LastName, professor.Email, professor.FacultyID).Scan(&professorID)
	if err != nil {
		return 0, err
	}
	return professorID, nil
}

func UpdateProfessor(pool *pgxpool.Pool, id int, professor models.Professor) error {
	_, err := pool.Exec(context.Background(), "UPDATE professor SET first_name=$1, last_name=$2, email=$3, faculty_id=$4 WHERE professor_id=$5", professor.FirstName, professor.LastName, professor.Email, professor.FacultyID, id)
	return err
}

func DeleteProfessor(pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM professor WHERE professor_id=$1", id)
	return err
}
