package postgres

import (
	"StudentManagement/storage"
	"fmt"
	"log"
)

const createSubjectQuery = `
INSERT INTO subjects(
class,
subject1
) VALUES (
	:class,
	:subject1
) returning *`

// const createSubjectQuery = `
// INSERT INTO subjects (
// 	class,
// 	subject1,
// 	subject2,
// 	subject3,
// 	subject4
// ) VALUES (
// 	:class,
// 	:subject1,
// 	:subject2,
// 	:subject3,
// 	:subject4)
// ON CONFLICT (class) DO UPDATE SET
//     subject1 = EXCLUDED.subject1,
//     subject2 = EXCLUDED.subject2,
//     subject3 = EXCLUDED.subject3,
//     subject4 = EXCLUDED.subject4
// 	returning * `

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

const listSubjectQuery = `SELECT * FROM subjects WHERE deleted_at IS NULL ORDER BY class ASC`

func (s PostgresStorage) ListSubject() ([]storage.Subject, error) {

	var subject []storage.Subject
	if err := s.DB.Select(&subject, listSubjectQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return subject, nil
}

const deleteSubjectByIdQuery = `UPDATE subjects SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteSubjectByID(id string) error {
	res, err := s.DB.Exec(deleteSubjectByIdQuery, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete subjects")
	}

	return nil
}

const getSubjectByIDQuery = `SELECT * FROM subjects WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetSubjectByID(id string) (*storage.Subject, error) {
	var u storage.Subject
	if err := s.DB.Get(&u, getSubjectByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

const UpdateSubjectQ = `
	UPDATE subjects SET
	class =:class,
	subject1 =:subject1
		WHERE id= :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateSubject(u storage.Subject) (*storage.Subject, error) {
	stmt, err := s.DB.PrepareNamed(UpdateSubjectQ)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &u, nil
}
