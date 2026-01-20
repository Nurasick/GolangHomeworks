package repository

import (
	"context"
	"errors"
	"time"
	"university/model"
)

type AttendanceRepositoryInterface interface {
	MarkAttendance(attendance *model.Attendance) error
	GetAttendanceByStudentID(studentID int) ([]model.Attendance, error)
	GetAttendanceByGroupIDandDate(groupID int, date time.Time) ([]model.Attendance, error)
}

func (r *UserRepository) MarkAttendance(attendance *model.Attendance) error {
	query := `
	insert into attendance(student_id, subject_id, visit_day, visited)
	values($1, $2, $3, $4);
	`

	err := r.Conn.QueryRow(context.Background(), query, attendance.StudentID, attendance.SubjectID, attendance.VisitDay, attendance.Visited).Scan()
	if err != nil {
		return errors.New("Failed to mark attendance: " + err.Error())
	}
	return nil
}

func (r *UserRepository) GetAttendanceByStudentID(studentID int) ([]model.Attendance, error) {
	query := `
	select a.id, a.student_id, s.name, a.subject_id, sub.name, a.visit_day, a.visited from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.student_id = $1;
	`
	rows, err := r.Conn.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, errors.New("Failed to get attendance by student ID: " + err.Error())
	}
	defer rows.Close()

	var attendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.StudentName, &attendance.SubjectID, &attendance.SubjectName, &attendance.VisitDay, &attendance.Visited)
		if err != nil {
			return nil, errors.New("Failed to scan attendance row: " + err.Error())
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}
func (r *UserRepository) GetAttendanceByGroupIDandDate(groupID int, date time.Time) ([]model.Attendance, error) {
	query := `
	select id, student_id, subject_id, visit_day, visited from attendance where group_id = $1 and visit_day = $2;
	`
	rows, err := r.Conn.Query(context.Background(), query, groupID, date)
	if err != nil {
		return nil, errors.New("Failed to get attendance by student ID: " + err.Error())
	}
	defer rows.Close()

	var attendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.SubjectID, &attendance.VisitDay, &attendance.Visited)
		if err != nil {
			return nil, errors.New("Failed to scan attendance row: " + err.Error())
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}
