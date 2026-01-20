package model

import "time"

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	GroupID   int       `json:"group_id"`
	BirthDate time.Time `json:"birth_date"`
	Year      int       `json:"year"`
	Gender    string    `json:"gender"`
	UserId    int       `json:"user_id"`
}
