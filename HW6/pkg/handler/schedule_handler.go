package handler

import (
	"net/http"
	"strconv"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	ScheduleService service.ScheduleServiceInterface
}

func NewScheduleHandler(scheduleService service.ScheduleServiceInterface) *ScheduleHandler {
	return &ScheduleHandler{ScheduleService: scheduleService}
}

// @Summary Get group schedule
// @Description Retrieve schedule for a specific group
// @Tags Schedules
// @Produce json
// @Security Bearer
// @Param id path int true "Group ID"
// @Success 200 {array} model.Schedule
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /schedule/group/{id} [get]
func (h *ScheduleHandler) GetScheduleByGroupID(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}

	data, err := h.ScheduleService.GetSchedulesByGroupID(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Get all schedule
// @Description Retrieve all schedules
// @Tags Schedules
// @Produce json
// @Security Bearer
// @Success 200 {array} model.Schedule
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /schedule/all_class_schedule [get]
func (h *ScheduleHandler) GetAllSchedules(c echo.Context) error {

	data, err := h.ScheduleService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
