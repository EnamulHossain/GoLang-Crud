package postgres

import (
	"StudentManagement/storage"
	"fmt"
	"log"

)

const listClassQuery = `SELECT * FROM classes WHERE deleted_at IS NULL ORDER BY id ASC`

func (s PostgresStorage) ListClass() ([]storage.Class, error) {

	var class []storage.Class
	if err := s.DB.Select(&class, listClassQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return class, nil
}

const createClassQuery = `
INSERT INTO classes (
 class_name
 ) VALUES (
	:class_name
	)
	returning *`

func (s PostgresStorage) CreateClass(c storage.Class) (*storage.Class, error) {

	var class storage.Class
	stmt, err := s.DB.PrepareNamed(createClassQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&class,c); err != nil {
		return nil, err
	}
	if class.ID == 0 {
		return nil, fmt.Errorf("unable to insert")
	}
	return &class, nil
}


const UpdateClassQ = `
	UPDATE classes SET
	    class_name =:class_name
		WHERE id= :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateClass(c storage.Class) (*storage.Class, error) {
	stmt, err := s.DB.PrepareNamed(UpdateClassQ)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&c, c); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &c, nil
}



const getClassByIDQuery = `SELECT * FROM classes WHERE id=$1 AND deleted_at IS NULL`

func (a PostgresStorage) GetClassByID(id string) (*storage.Class, error) {
	var s storage.Class
	if err := a.DB.Get(&s, getClassByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}


// const deleteClassByIdQuery = `UPDATE classes SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

const deleteClassByIdQuery = `DELETE FROM classes WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteClassByID(id string) error {

	res, err := s.DB.Exec(deleteClassByIdQuery, id)
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
		return fmt.Errorf("unable to delete class")
	}

	return nil
}