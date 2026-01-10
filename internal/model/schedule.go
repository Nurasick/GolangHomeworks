package model

type Schedule struct {
	ID        int    `json:"id"`
	GroupName string `json:"group_name"`
	ClassName string `json:"class_name"`
	Day       string `json:"day_of_week"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
}
