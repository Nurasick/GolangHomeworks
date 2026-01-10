package handler

import (
	"context"
	"net/http"
	"strconv"

	"university/internal/db"
	"university/internal/model"

	"github.com/labstack/echo/v4"
)

func GetAllSchedule(c echo.Context) error {
	query := `
	SELECT sc.id, g.name, sc.class_name, sc.day_of_week, sc.starts_at, sc.ends_at
	FROM class_schedule sc
	JOIN groups g ON sc.group_id = g.id
	ORDER BY sc.day_of_week;
	`

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedule"})
	}
	defer rows.Close()

	var result []model.Schedule

	for rows.Next() {
		var r model.Schedule
		rows.Scan(&r.ID, &r.GroupName, &r.ClassName, &r.Day, &r.StartsAt, &r.EndsAt)
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
	SELECT sc.id, g.name, sc.class_name, sc.day_of_week, sc.starts_at, sc.ends_at
	FROM class_schedule sc
	JOIN groups g ON sc.group_id = g.id
	Where g.id = $1
	ORDER BY sc.day_of_week;
	`

	rows, err := db.Conn.Query(context.Background(), query, groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedule"})
	}
	defer rows.Close()

	var result []model.Schedule

	for rows.Next() {
		var r model.Schedule
		rows.Scan(&r.ID, &r.GroupName, &r.ClassName, &r.Day, &r.StartsAt, &r.EndsAt)
		result = append(result, r)
	}
	return c.JSON(http.StatusOK, result)
}
