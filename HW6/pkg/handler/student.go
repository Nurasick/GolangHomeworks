package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"university/helpers/db"
	"university/model"

	"github.com/labstack/echo/v4"
)

func GetStudentByID(c echo.Context) error {
	idParam := c.Param("id")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}
	query := `
	SELECT s.id, s.name,s.birth_date, g.name,gender.name
	FROM students s
	JOIN groups g ON s.group_id=g.id
	LEFT JOIN genders gender ON gender.id = s.gender_id
	WHERE s.id =$1;
	`

	var stud model.Student
	var birth time.Time

	err = db.Conn.QueryRow(context.Background(), query, idParam).Scan(&stud.ID, &stud.Name, &birth, &stud.GroupName, &stud.Gender)

	if err != nil {
		log.Println("QUERY ERROR:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve student"})
	}
	stud.BirthDate = birth.Format("2006-01-02")
	return c.JSON(http.StatusOK, stud)
}
