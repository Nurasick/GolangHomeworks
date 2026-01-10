package main

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GroupID int    `json:"group_id"`
}

var students = []Student{
	{ID: 1, Name: "Aigerim", GroupID: 1},
	{ID: 2, Name: "Asel", GroupID: 2},
	{ID: 3, Name: "Dana", GroupID: 3},
}

func main() {
	e := echo.New()
	e.GET("/students", getStudents)
	e.GET("/students/:id", getStudentByID)
	e.POST("/students", createStudent)
	e.GET("/groups", getGroups)

	e.Logger.Fatal(e.Start(":8080"))

}

func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}
func getStudentByID(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid id",
		})
	}
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Student not found",
	})
}

func createStudent(c echo.Context) error {
	var newStudent Student

	if err := c.Bind(&newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}
	newStudent.ID = len(students) + 1

	students = append(students, newStudent)

	return c.JSON(http.StatusCreated, newStudent)
}
func getGroups(c echo.Context) error {
	groups := []string{"ENG-101", "HUM-201"}
	return c.JSON(http.StatusOK, groups)
}
