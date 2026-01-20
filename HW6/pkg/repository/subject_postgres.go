package repository

import (
	"context"
	"errors"
	"university/model"
)

type SubjectRepositoryInterface interface {
	CreateSubject(subject *model.Subject) error
	GetSubjectByID(id int) (*model.Subject, error)
	GetAllSubject() ([]model.Subject, error)
}

func (r *UserRepository) CreateSubject(subject *model.Subject) (int, error) {
	var id int
	query := `
	insert into subjects(name)
	values($1) returning id
	`
	err := r.Conn.QueryRow(context.Background(), query, subject.Name).Scan(&id)
	if err != nil {
		return 0, errors.New("Failed to create a subject: " + err.Error())
	}

	return id, nil
}

func (r *UserRepository) GetSubjectByID(id int) (model.Subject, error) {
	query := `select id, name from subjects where id=$1;`

	var subject model.Subject
	err := r.Conn.QueryRow(context.Background(), query, id).Scan(
		&subject.ID, &subject.Name,
	)
	if err != nil {
		return model.Subject{}, errors.New("Failed to get a subject by ID: " + err.Error())
	}
	return subject, nil
}

func (r *UserRepository) GetAllSubjects() ([]model.Subject, error) {
	query := `select id, name from subjects;`

	var subjects []model.Subject
	rows, err := r.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve subjects: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var subject model.Subject
		err := rows.Scan(&subject.ID, &subject.Name)
		if err != nil {
			return nil, errors.New("Failed to scan subjects " + err.Error())
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}
