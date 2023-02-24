package postgres

import (
	"StudentManagement/storage"
	"fmt"
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



// edit


const GetMarkEditQuery = `SELECT student_subjects.id,marks FROM student_subjects WHERE student_subjects.id = $1 AND deleted_at IS NULL;`

func (s PostgresStorage) MarkEdit(id string) (*storage.MarkEdit, error) {
	var u storage.MarkEdit
	if err := s.DB.Get(&u, GetMarkEditQuery, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &u, nil
}




const UpdateMarksbyIDQuery = `UPDATE student_subjects
SET marks = $1
WHERE id = $2 AND deleted_at IS NULL;`

func (s PostgresStorage) UpdateMarksbyID(marks string,id string) error {
	res, err := s.DB.Exec(UpdateMarksbyIDQuery, marks,id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return nil
	}

	return nil
}