package service

import (
	"errors"
	"university/model"
)

func (r *UserService) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	userID := teacher.UserId
	user, err := r.Repo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("User does not exist: " + err.Error())
	}
	if user.RoleID != 2 {
		return nil, errors.New("User is not a teacher")
	}
	teacherID, err := r.Repo.CreateTeacher(teacher)
	if err != nil {
		return nil, errors.New("Failed to create a teacher: " + err.Error())
	}
	teacher.ID = teacherID
	return teacher, nil
}

func (r *UserService) GetScheduleByTeacherID(id int) ([]model.Schedule, error) {
	teacher, err := r.Repo.GetTeacherByID(id)
	if err != nil {
		return nil, errors.New("Failed to get teacher: " + err.Error())
	}
	schedules, err := r.Repo.GetScheduleByTeacherID(teacher.ID)
	if err != nil {
		return nil, errors.New("Failed to get schedule by teacher ID: " + err.Error())
	}
	return schedules, nil
}
