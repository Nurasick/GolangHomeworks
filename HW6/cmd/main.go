package main

import (
	"context"
	"university/helpers/db"
	"university/pkg/handler"
	"university/pkg/middleware"
	"university/pkg/repository"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Connect("postgres://postgres:1234@localhost:5432/university")
	defer db.Conn.Close(context.Background())

	// Initialize services
	UserRepository := &repository.UserRepository{db.Conn}
	AuthService := &service.AuthService{Repo: UserRepository}
	UserService := &service.UserService{Repo: UserRepository}

	// Initialize handlers
	AuthHandler := handler.NewAuthHandler(AuthService)
	UserHandler := handler.NewUserHandler(UserService)

	e := echo.New()

	e.GET("/students/:id", handler.GetStudentByID)
	e.GET("/all_class_schedule", handler.GetAllSchedule)
	e.GET("/schedule/group/:id", handler.GetScheduleByGroupID)

	e.POST("/attendance/subject", handler.PostAttendance)
	e.GET("/attendanceBySubjectID/:id", handler.GetAttendanceBySubjectID)
	e.GET("/attendanceByStudentID/:id", handler.GetAttendanceByStudentID)

	auth := e.Group("/api/auth")
	auth.POST("/register", AuthHandler.Register)
	auth.POST("/login", AuthHandler.Login)

	users := e.Group("/api/users", middleware.JWTAuth)
	users.GET("/me", UserHandler.Me)

	e.Logger.Fatal(e.Start(":8080"))
}
