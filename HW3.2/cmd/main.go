package main

import (
	"university/internal/db"
	"university/internal/handler"

	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Connect("postgres://postgres:1234@localhost:5432/university")
	defer db.Conn.Close(context.Background())

	e := echo.New()

	e.GET("/students/:id", handler.GetStudentByID)
	e.GET("/all_class_schedule", handler.GetAllSchedule)
	e.GET("/schedule/group/:id", handler.GetScheduleByGroupID)

	e.Logger.Fatal(e.Start(":8080"))
}
