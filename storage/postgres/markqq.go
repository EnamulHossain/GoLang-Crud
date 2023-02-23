package postgres

import (
	"StudentManagement/storage"
	"log"
)

const markd = `SELECT subjects.subject1, subjects.class, students.first_name, students.last_name, students.roll,student_subjects.subject_id,student_subjects.id
FROM subjects
FULL OUTER JOIN student_subjects ON subjects.id = student_subjects.subject_id
FULL OUTER JOIN students ON students.id = student_subjects.student_id
WHERE students.id = $1
ORDER BY subjects.subject1;`

func (s PostgresStorage) GetMarkInputOptionByID(id string) ([]storage.MarkInputStore, error) {
	var u []storage.MarkInputStore
	if err := s.DB.Select(&u, markd, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}

const createmarkQuery = `
UPDATE student_subjects
SET marks = :marks
WHERE id = :id
	returning *;`

func (p PostgresStorage) Markcreate(s storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, _ := p.DB.PrepareNamed(createmarkQuery)

	stmt.Get(&s, s)

	return &s, nil

}
