package model

import "time"

type Attendance struct {
	ID          int       `json:"id"`
	StudentID   int       `json:"student_id"`
	StudentName string    `json:"student_name"`
	SubjectID   int       `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	VisitDay    time.Time `json:"visit_day"`
	Visited     bool      `json:"visited"`
}
type AttendanceResponse struct {
	ID          int       `json:"id"`
	StudentName string    `json:"student_name"`
	SubjectName string    `json:"subject_name"`
	VisitDay    time.Time `json:"visit_day"`
	Visited     bool      `json:"visited"`
}
