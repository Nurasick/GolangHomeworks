package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
	"university/internal/db"
	"university/internal/model" //importing db and models

	"github.com/labstack/echo/v4"
)

func PostAttendance(c echo.Context) error { //POST attendance handler
	var att model.Attendance

	if err := c.Bind(&att); err != nil {
		log.Println("BIND ERROR:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if _, err := time.Parse("2006-01-02", att.VisitDay); err != nil {
		log.Println("DATE PARSE ERROR:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format. Use YYYY-MM-DD"})
	}
	_, err := db.Conn.Exec(context.Background(), `
		insert into attendance (student_id, subject_id, visit_day, visited)
		values ($1, $2, $3, $4);
	`, att.StudentID, att.SubjectID, att.VisitDay, att.Visited)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Attendance recorded successfully"})
}

func GetAttendanceBySubjectID(c echo.Context) error { //GET attendance by subject ID handler
	subjectID := c.Param("id")

	if _, err := strconv.Atoi(subjectID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	query := `
	select a.id, a.student_id, s.name as student_name, a.subject_id, sub.name as subject_name, a.visit_day, a.visited
	from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.subject_id = $1
	order by a.visit_day;
	`

	rows, err := db.Conn.Query(context.Background(), query, subjectID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve attendance"})
	}

	defer rows.Close()

	var result []model.AttendanceResponse
	for rows.Next() {
		var a model.Attendance
		var visitDay time.Time
		if err := rows.Scan(&a.ID, &a.StudentID, &a.StudentName, &a.SubjectID, &a.SubjectName, &visitDay, &a.Visited); err != nil {
			log.Println("SCAN ERROR:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse attendance data"})
		}
		a.VisitDay = visitDay.Format("2006-01-02")

		resp := model.AttendanceResponse{
			ID:          a.ID,
			StudentName: a.StudentName,
			SubjectName: a.SubjectName,
			VisitDay:    a.VisitDay,
			Visited:     a.Visited,
		}

		result = append(result, resp)
	}
	return c.JSON(http.StatusOK, result)
}

func GetAttendanceByStudentID(c echo.Context) error {
	studentID := c.Param("id")

	if _, err := strconv.Atoi(studentID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	query := `
	select a.id, a.student_id, s.name as student_name, a.subject_id, sub.name as subject_name, a.visit_day, a.visited
	from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.student_id = $1
	order by a.visit_day;
	`

	rows, err := db.Conn.Query(context.Background(), query, studentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve attendance"})
	}

	defer rows.Close()

	var result []model.AttendanceResponse
	for rows.Next() {
		var a model.Attendance
		var visitDay time.Time
		if err := rows.Scan(&a.ID, &a.StudentID, &a.StudentName, &a.SubjectID, &a.SubjectName, &visitDay, &a.Visited); err != nil {
			log.Println("SCAN ERROR:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse attendance data"})
		}
		a.VisitDay = visitDay.Format("2006-01-12")

		resp := model.AttendanceResponse{
			ID:          a.ID,
			StudentName: a.StudentName,
			SubjectName: a.SubjectName,
			VisitDay:    a.VisitDay,
			Visited:     a.Visited,
		}

		result = append(result, resp)
	}
	return c.JSON(http.StatusOK, result)

}
