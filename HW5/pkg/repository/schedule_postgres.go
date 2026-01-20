package repository

import (
	"context"
	"errors"
	"university/model"
)

type ScheduleRepositoryInterface interface {
	CreateSchedule(*model.Schedule) (int, error)
	GetScheduleByGroupID(groupID int) ([]model.Schedule, error)
	GetScheduleByTeacherID(teacherID int) ([]model.Schedule, error)
	DeleteSchedule(id int) error
}

func (r *UserRepository) CreateSchedule(schedule *model.Schedule) (int, error) {
	query := `
	insert into schedule(group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id)
	values($1, $2, $3, $4, $5, $6);
	`
	var id int

	err := r.Conn.QueryRow(context.Background(), query, schedule.GroupID, schedule.SubjectID, schedule.Day, schedule.StartsAt, schedule.EndsAt, schedule.TeacherID).Scan(&id)

	if err != nil {
		return 0, errors.New("Failed to create schedule: " + err.Error())
	}
	return id, nil
}

func (r *UserRepository) GetScheduleByGroupID(groupID int) ([]model.Schedule, error) {
	query := `
	select id, group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id from class_schedule where group_id = $1;
	`
	rows, err := r.Conn.Query(context.Background(), query, groupID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedule by group ID: " + err.Error())
	}
	defer rows.Close()
	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.Day, &schedule.StartsAt, &schedule.EndsAt, &schedule.TeacherID)
		if err != nil {
			return nil, errors.New("Failed to scan schedule by group ID: " + err.Error())
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (r *UserRepository) GetScheduleByTeacherID(teacherID int) ([]model.Schedule, error) {
	query := `
	select id, group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id from class_schedule where teacher_id = $1;
	`
	rows, err := r.Conn.Query(context.Background(), query, teacherID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedule by teacher ID: " + err.Error())
	}
	defer rows.Close()
	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.Day, &schedule.StartsAt, &schedule.EndsAt, &schedule.TeacherID)
		if err != nil {
			return nil, errors.New("Failed to scan schedule by teacher ID: " + err.Error())
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (r *UserRepository) DeleteSchedule(id int) error {
	query := `
	delete from class_schedule where id = $1;
	`
	_, err := r.Conn.Exec(context.Background(), query, id)
	if err != nil {
		return errors.New("Failed to delete schedule: " + err.Error())
	}
	return nil
}
