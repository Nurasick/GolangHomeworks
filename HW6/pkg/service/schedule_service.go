package service

import (
	"errors"
	"university/model"
)

func (r *UserService) CreateSchedule(schedule *model.Schedule) (*model.Schedule, error) {
	if schedule.Day < 0 || schedule.Day > 7 {
		return nil, errors.New("Invalid day of the week")
	}
	scheduleID, err := r.Repo.CreateSchedule(schedule)
	if err != nil {
		return nil, errors.New("Failed to create schedule: " + err.Error())
	}
	schedule.ID = scheduleID
	return schedule, nil
}

func (r *UserService) GetSchedulesbyGroupID(groupID int) ([]model.Schedule, error) {

}
