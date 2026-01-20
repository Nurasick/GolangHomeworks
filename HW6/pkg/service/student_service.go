package service

import (
	"errors"
	"university/model"
)

func (r *UserService) CreateStudent(student *model.Student) (*model.Student, error) {
	userID := student.UserId
	user, err := r.Repo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("User does not exist: " + err.Error())
	}
	if user.RoleID != 3 {
		return nil, errors.New("User is not a student: ")
	}
	studentID, err := r.Repo.CreateStudent(student)
	if err != nil {
		return nil, errors.New("Failed to create student: " + err.Error())
	}
	student.ID = studentID
	return student, nil
}

func (r *UserService) GetStudentByID(id int) (*model.Student, error) {
	student, err := r.Repo.GetStudentByID(id)
	if err != nil {
		return nil, errors.New("Failed to get student: " + err.Error())
	}
	return &student, nil
}

func (r *UserService) GetStudentByUserID(id int) (*model.Student, error) {
	student, err := r.Repo.GetStudentByUserID(id)
	if err != nil {
		return nil, errors.New("Failed to get student: " + err.Error())
	}
	return &student, nil
}

func (r *UserService) GetStudentAttendance(studentID int) ([]model.AttendanceResponse, error) {
	attendances, err := r.Repo.GetAttendanceByStudentID(studentID)
	if err != nil {
		return nil, errors.New("Failed to get student attendance: " + err.Error())
	}
	var attendanceResponses []model.AttendanceResponse
	for _, attendance := range attendances {
		attendanceResponses = append(attendanceResponses, model.AttendanceResponse{
			ID:          attendance.ID,
			StudentName: attendance.StudentName,
			SubjectName: attendance.SubjectName,
			VisitDay:    attendance.VisitDay,
			Visited:     attendance.Visited,
		})
	}
	return attendanceResponses, nil
}
