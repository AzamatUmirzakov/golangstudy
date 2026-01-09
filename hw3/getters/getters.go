package getters

import (
	"context"
	"hw3/types"

	"github.com/jackc/pgx/v5"
)

func GetStudentByID(connection *pgx.Conn, id int) (types.Student, error) {
	var student types.Student
	err := connection.QueryRow(context.Background(), "SELECT * FROM student WHERE student_id=$1", id).Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.Email, &student.Gender, &student.BirthDate, &student.GroupID)

	if err != nil {
		return types.Student{}, err
	}

	return student, nil
}

func GetGroupByID(connection *pgx.Conn, groupId int) (types.StudentGroup, error) {
	var group types.StudentGroup
	err := connection.QueryRow(context.Background(), "SELECT * FROM student_group WHERE group_id = $1", groupId).Scan(&group.GroupID, &group.FacultyID, &group.GroupName)

	if err != nil {
		return types.StudentGroup{}, err
	}

	return group, nil
}

func GetTimetables(connection *pgx.Conn) ([]types.Timetable, error) {
	rows, err := connection.Query(context.Background(), "SELECT * FROM timetable ORDER BY start_time ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// gathering data into a slice
	var timetables []types.Timetable
	for rows.Next() {
		var timetable types.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.Subject, &timetable.Professor)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}

	return timetables, nil
}

func GetTimetableByGroupID(connection *pgx.Conn, groupId int) ([]types.Timetable, error) {
	var timetables []types.Timetable
	rows, err := connection.Query(context.Background(), "SELECT * FROM timetable WHERE group_id = $1 ORDER BY start_time ASC", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var timetable types.Timetable
		err := rows.Scan(&timetable.TimetableID, &timetable.FacultyID, &timetable.GroupID, &timetable.StartTime, &timetable.EndTime, &timetable.Weekday, &timetable.Location, &timetable.Subject, &timetable.Professor)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, timetable)
	}
	return timetables, nil
}
