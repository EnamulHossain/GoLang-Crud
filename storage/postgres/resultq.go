package postgres

import (
	"StudentManagement/storage"
	"log"
)

const res = `SELECT students.first_name, students.last_name, subjects.class, students.roll, subjects.subject1,student_subjects.marks
FROM subjects
FULL OUTER JOIN student_subjects ON subjects.id = student_subjects.subject_id
FULL OUTER JOIN students ON students.id = student_subjects.student_id;`

func (s PostgresStorage) Resul() ([]storage.Result, error) {
	var resultt []storage.Result
	if err := s.DB.Select(&resultt, res); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resultt, nil
}
