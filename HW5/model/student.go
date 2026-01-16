package model

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"group_name"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}
