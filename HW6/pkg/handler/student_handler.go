package handler

import (
	"net/http"
	"strconv"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	StudentService service.StudentServiceInterface
}

// NewUserHandler creates a new instance of UserHandler
func NewStudentHandler(studentService service.StudentServiceInterface) *StudentHandler {
	return &StudentHandler{StudentService: studentService}
}

// @Summary Get student attendance
// @Description Get attendance records for the current student
// @Tags Students
// @Produce json
// @Security Bearer
// @Param id path int true "Student ID"
// @Success 200 {array} model.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student/{id}/attendance [get]
func (h *StudentHandler) MyAttendance(c echo.Context) error {
	userID := c.Get("userID").(int)

	stud, err := h.StudentService.GetStudentByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	data, err := h.StudentService.GetStudentAttendance(stud.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Get student by ID
// @Description Retrieve student information by user ID
// @Tags Students
// @Produce json
// @Security Bearer
// @Param id path int true "User ID"
// @Success 200 {object} model.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student/{id} [get]
func (h *StudentHandler) GetStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}

	stud, err := h.StudentService.GetStudentByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	data := model.StudentResponse{
		ID:        stud.ID,
		Name:      stud.Name,
		BirthDate: stud.BirthDate,
		GroupID:   stud.GroupID,
		GroupName: stud.GroupName,
		Gender:    stud.Gender,
	}

	return c.JSON(http.StatusOK, data)
}
