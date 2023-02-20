package postgres

import (
	"StudentManagement/storage"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const listStudentQuery = `SELECT * FROM students WHERE deleted_at IS NULL ORDER BY id ASC`

func (s PostgresStorage) ListStudent() ([]storage.Student, error) {

	var student []storage.Student
	if err := s.DB.Select(&student, listStudentQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return student, nil
}

const createStudentQuery = `
INSERT INTO students (
 first_name, 
 last_name, 
 class,
 roll, 
 email, 
 password
 ) VALUES (
	:first_name, 
	:last_name, 
	:class, 
	:roll, 
	:email, 
	:password
	)
	returning *`

func (s PostgresStorage) CreateStudent(u storage.Student) (*storage.Student, error) {

	var student storage.Student
	stmt, err := s.DB.PrepareNamed(createStudentQuery)
	if err != nil {
		log.Fatalln(err)
	}
	// HAsh
	HassPass, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(HassPass)
	//ENd HAsh


	if err := stmt.Get(&student, u); err != nil {
		return nil, err
	}
	if student.ID == 0 {
		return nil, fmt.Errorf("unable to insert")
	}
	return &student, nil
}


const UpdateStudentQ = `
	UPDATE students SET
	    first_name =:first_name,
		last_name =:last_name,
		roll = :roll,
		email = :email,
		password = :password
		WHERE id= :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateStudent(u storage.Student) (*storage.Student, error) {
	stmt, err := s.DB.PrepareNamed(UpdateStudentQ)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec(u)
	if err != nil {
		log.Fatalln(err)
	}
	Rcount, err := res.RowsAffected()
	if Rcount < 1 || err != nil {
		log.Fatalln(err)
	}
	return &u, nil
}

const getStudentByIDQuery = `SELECT * FROM students WHERE id=$1 AND deleted_at IS NULL`


func (s PostgresStorage) GetStudentByID(id string) (*storage.Student, error) {
	var u storage.Student
	if err := s.DB.Get(&u, getStudentByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}



const getStudentByUsernameQuery = `SELECT * FROM students WHERE username=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetStudentByUsername(username string) (*storage.Student, error) {
	var u storage.Student
	if err := s.DB.Get(&u, getStudentByUsernameQuery, username); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}



const deleteStudentByIdQuery = `UPDATE students SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteStudentByID(id string) error {
	res, err := s.DB.Exec(deleteStudentByIdQuery, id)
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
		return fmt.Errorf("unable to delete student")
	}

	return nil
}