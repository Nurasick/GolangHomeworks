package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"university/internal/db"
	"university/internal/model"

	"github.com/labstack/echo/v4"
)

func GetAllSchedule(c echo.Context) error {
	query := `
	select cs.id, g.name as group_name, sub.name as subject_name, cs.day_of_week, cs.starts_at, cs.ends_at
	from class_schedule cs
	join groups g on cs.group_id = g.id
	join subjects sub on cs.subject_id = sub.id
	order by cs.day_of_week;
	`

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		log.Println("QUERY ERROR:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedule"})
	}
	defer rows.Close()

	var result []model.Schedule

	for rows.Next() {
		var r model.Schedule
		var starts, ends time.Time
		err := rows.Scan(&r.ID, &r.GroupName, &r.SubjectName, &r.Day, &starts, &ends)
		if err != nil {
			log.Println("SCAN ERROR:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse schedule data"})
		}
		r.StartsAt = starts.Format("15:04")
		r.EndsAt = ends.Format("15:04")
		result = append(result, r)
	}
	return c.JSON(http.StatusOK, result)
}

func GetScheduleByGroupID(c echo.Context) error {

	groupID := c.Param("id")
	_, err := strconv.Atoi(groupID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid group ID"})
	}

	query := `
	select cs.id, g.name as group_name, sub.name as subject_name, cs.day_of_week, cs.starts_at, cs.ends_at
	from class_schedule cs
	join groups g on cs.group_id = g.id
	join subjects sub on cs.subject_id = sub.id
	where g.id = $1
	order by cs.day_of_week;
	`

	rows, err := db.Conn.Query(context.Background(), query, groupID)
	if err != nil {
		log.Println("QUERY ERROR:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedule"})
	}
	defer rows.Close()

	var result []model.Schedule

	for rows.Next() {
		var r model.Schedule
		var starts, ends time.Time
		err := rows.Scan(&r.ID, &r.GroupName, &r.SubjectName, &r.Day, &starts, &ends)
		if err != nil {
			log.Println("SCAN ERROR:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse schedule data"})
		}
		r.StartsAt = starts.Format("15:04")
		r.EndsAt = ends.Format("15:04")
		result = append(result, r)
	}
	return c.JSON(http.StatusOK, result)
}
