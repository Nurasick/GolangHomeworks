package service

import (
	"errors"
	"university/model"
	"university/pkg/repository"
)

type ScheduleServiceInterface interface {
	CreateSchedule(schedule *model.Schedule) (*model.Schedule, error)
	GetSchedulesByGroupID(groupID int) ([]model.Schedule, error)
	GetAllSchedule() ([]model.Schedule, error)
}

type ScheduleService struct {
	ScheduleRepo repository.ScheduleRepositoryInterface
}

func NewScheduleService(
	scheduleRepo *repository.ScheduleRepository,
) *ScheduleService {
	return &ScheduleService{
		ScheduleRepo: scheduleRepo,
	}
}

func (r *ScheduleService) CreateSchedule(schedule *model.Schedule) (*model.Schedule, error) {
	if schedule.Day < 0 || schedule.Day > 7 {
		return nil, errors.New("Invalid day of the week")
	}
	scheduleID, err := r.ScheduleRepo.CreateSchedule(schedule)
	if err != nil {
		return nil, errors.New("Failed to create schedule: " + err.Error())
	}
	schedule.ID = scheduleID
	return schedule, nil
}

func (r *ScheduleService) GetSchedulesByGroupID(groupID int) ([]model.Schedule, error) {
	schedules, err := r.ScheduleRepo.GetScheduleByGroupID(groupID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedules by groupID: " + err.Error())
	}
	return schedules, nil

}
func (r *ScheduleService) GetAllSchedule() ([]model.Schedule, error) {
	schedules, err := r.ScheduleRepo.GetAllSchedule()

	if err != nil {
		return nil, errors.New("Failed to retrieve schedules: " + err.Error())
	}
	return schedules, nil
}
