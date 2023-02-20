package postgres

import (
	"StudentManagement/storage"
	"log"
)

const getSubjectByClassIDQuery = `SELECT * FROM subjects WHERE class=$1`

func (s PostgresStorage) GetSubjectByClassID(class int) ([]storage.Subject, error) {

	var u []storage.Subject
	if err := s.DB.Select(&u, getSubjectByClassIDQuery, class); err != nil {
		log.Println(err)
		return nil, err
	}
	return u, nil
}

const insertMarkQuery = `
	INSERT INTO student_subjects(
		student_id,
		subject_id,
        marks
		)  
	VALUES(
		:student_id,
		:subject_id,
		:marks
		)RETURNING *;
	`

func (p PostgresStorage) InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, err := p.DB.PrepareNamed(insertMarkQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}
