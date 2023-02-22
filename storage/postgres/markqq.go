package postgres

import (
	"StudentManagement/storage"
	"log"
)


const markd =`SELECT subjects.subject1, subjects.class, students.first_name, students.last_name, students.roll
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