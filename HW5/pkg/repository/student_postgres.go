package repository

import (
	"context"
	"errors"
	"university/model"
)

type StudentRepositoryInterface interface {
	CreateStudent(*model.Student) (int, error)
	GetStudentByID(id int) (model.Student, error)
	GetStudentByUserID(userID int) (model.Student, error)
	GetAllStudents() ([]model.Student, error)
	UpdateStudent(*model.Student) error
}

func (r *UserRepository) CreateStudent(student *model.Student) (int, error) {
	query := `insert into students (name,birth_date,group_id,gender_id,user_id,year) 
	values ($1,$2,$3,$4,$5,$6) returning id;`

	var id int
	err := r.Conn.QueryRow(context.Background(), query, student.Name, student.BirthDate, student.GroupID, student.Gender, student.UserId, student.Year).Scan(&id)
	if err != nil {
		return 0, errors.New("Failed to create student: " + err.Error())
	}
	student.ID = id
	return id, nil

}

func (r *UserRepository) GetStudentByID(id int) (model.Student, error) {
	query := `
	SELECT s.id, s.name, s.birth_date, s.group_id, g.gender, s.user_id, s.year
	FROM students s
	JOIN genders g ON g.id = s.gender_id
	WHERE s.id =$1;`

	var student model.Student

	err := r.Conn.QueryRow(context.Background(), query, id).Scan(&student.ID, &student.Name, &student.BirthDate, &student.GroupID, &student.Gender, &student.UserId, &student.Year)
	if err != nil {
		return model.Student{}, errors.New("Failed to retrieve student: " + err.Error())
	}
	return student, nil
}

func (r *UserRepository) GetStudentByUserID(userID int) (model.Student, error) {
	query := `
	SELECT s.id, s.name, s.birth_date, s.group_id, g.gender, s.user_id, s.year
	FROM students s
	JOIN genders g ON g.id = s.gender_id
	WHERE s.user_id =$1;`

	var student model.Student

	err := r.Conn.QueryRow(context.Background(), query, userID).Scan(&student.ID, &student.Name, &student.BirthDate, &student.GroupID, &student.Gender, &student.UserId, &student.Year)
	if err != nil {
		return model.Student{}, errors.New("Failed to retrieve student: " + err.Error())
	}
	return student, nil
}
func (r *UserRepository) GetAllStudents() ([]model.Student, error) {
	query := `
	SELECT s.id, s.name, s.birth_date, s.group_id, g.gender, s.user_id, s.year
	FROM students s
	JOIN genders g ON g.id = s.gender_id;`

	rows, err := r.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve students: " + err.Error())
	}
	defer rows.Close()
	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.BirthDate, &student.GroupID, &student.Gender, &student.UserId, &student.Year)
		if err != nil {
			return nil, errors.New("Failed to scan student: " + err.Error())
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *UserRepository) UpdateStudent(student *model.Student) error {
	query := `UPDATE students SET name=$1, birth_date=$2, group_id=$3, gender_id=$4, year=$5 where id = $6;`

	err := r.Conn.QueryRow(context.Background(), query, student.Name, student.BirthDate, student.GroupID, student.Gender, student.Year, student.ID).Scan()
	if err != nil {
		return errors.New("Failed to update student: " + err.Error())
	}
	return nil
}
