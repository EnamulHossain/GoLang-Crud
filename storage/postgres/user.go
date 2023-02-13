package postgres

import (
	"StudentManagement/storage"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

const listQuery = `SELECT * FROM users WHERE deleted_at IS NULL ORDER BY id ASC`

func (s PostgresStorage) ListUser() ([]storage.User, error) {

	var user []storage.User
	if err := s.DB.Select(&user, listQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}

const createUserQuery = `
	INSERT INTO users(
		name,
		email,
		password
		)  VALUES(
			:name,
		:email,
		:password
		)
		returning *`

func (s PostgresStorage) CreateUser(u storage.User) (*storage.User, error) {

	
	stmt, _ := s.DB.PrepareNamed(createUserQuery)
// HAsh
	HassPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(HassPass)

//ENd HAsh

	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to insert user into db")
	}
	return &u, nil
}

const UpdateQQ = `
UPDATE Users SET
	name =:name,
	email = :email,
	password = :password
	WHERE id= :id AND deleted_at IS NULL RETURNING *;
`

func (s PostgresStorage) UpdateUser(u storage.User) (*storage.User, error) {

	stmt, err := s.DB.PrepareNamed(UpdateQQ)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &u, nil
}

const getUserByIDQuery = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetUserByID(id string) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, getUserByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}


const getUserByUsernameQuery = `SELECT * FROM users WHERE name=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetUserByUsername(name string) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, getUserByUsernameQuery, name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}




const deleteUserByIdQuery = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteUserByID(id string) error {
	res, err := s.DB.Exec(deleteUserByIdQuery, id)
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
		return fmt.Errorf("unable to delete user")
	}

	return nil
}