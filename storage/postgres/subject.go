package postgres

import (
	"StudentManagement/storage"
	"fmt"
	"log"
)

// const createSubjectQuery = `
// INSERT INTO subjects(
// class,
// subject1,
// subject2,
// subject3,
// subject4
// ) VALUES (
// 	:class,
// 	:subject1,
// 	:subject2,
// 	:subject3,
// 	:subject4
// ) returning *`

const createSubjectQuery = `
INSERT INTO subjects (
	class, 
	subject1, 
	subject2, 
	subject3, 
	subject4
) VALUES ( 
	:class, 
	:subject1, 
	:subject2, 
	:subject3, 
	:subject4)
ON CONFLICT (class) DO UPDATE SET
    subject1 = EXCLUDED.subject1,
    subject2 = EXCLUDED.subject2,
    subject3 = EXCLUDED.subject3,
    subject4 = EXCLUDED.subject4
	returning * `

func (s PostgresStorage) CreateSubject(u storage.Subject) (*storage.Subject, error) {

	var subject storage.Subject
	stmt, err := s.DB.PrepareNamed(createSubjectQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&subject, u); err != nil {
		return nil, err
	}
	if subject.ID == 0 {
		return nil, fmt.Errorf("unable to insert")
	}
	return &subject, nil
}

const listSubjectQuery = `SELECT * FROM subjects WHERE deleted_at IS NULL ORDER BY id ASC`

func (s PostgresStorage) ListSubject() ([]storage.Subject, error) {

	var subject []storage.Subject
	if err := s.DB.Select(&subject, listSubjectQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return subject, nil
}
