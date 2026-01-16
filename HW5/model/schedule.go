package model

type Schedule struct {
	ID          int    `json:"id"`
	GroupName   string `json:"group_name"`
	SubjectName string `json:"subject_name"`
	Day         int    `json:"day_of_week"`
	StartsAt    string `json:"starts_at"`
	EndsAt      string `json:"ends_at"`
}
