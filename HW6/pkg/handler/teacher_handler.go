package handler

import (
	"net/http"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type TeacherHandler struct {
	TeacherService    service.TeacherServiceInterface
	AttendanceService service.AttendanceServiceInterface
}

func NewTeacherHandler(
	teacherService service.TeacherServiceInterface,
	attendanceService service.AttendanceServiceInterface,
) *TeacherHandler {
	return &TeacherHandler{
		TeacherService:    teacherService,
		AttendanceService: attendanceService,
	}
}

func (h *TeacherHandler) MarkAttendance(c echo.Context) error {
	var req model.Attendance
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}
	//roleID := c.Get("RoleID").(int)

	err := h.AttendanceService.MarkAttendance(req.StudentID, req.SubjectID, 1, req.VisitDay, req.Visited)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"data": "attendance saved"})
}
