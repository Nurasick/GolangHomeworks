package model

import "time"

type Schedule struct {
	ID        int       `json:"id"`
	GroupID   int       `json:"group_id"`
	SubjectID int       `json:"subject_id"`
	Day       int       `json:"day_of_week"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
	TeacherID int       `json:"teacher_id"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
