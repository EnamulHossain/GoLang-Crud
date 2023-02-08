package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)


type User struct {
	ID        int          `db:"id" form:"-"`
	Name      string       `db:"name" form:"name"`
	Email     string       `db:"email" form:"email"`
	Password  string       `db:"password" form:"password"`
	Status    bool         `db:"status" form:"status"`
	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
	// CSRFToken string       `db:"-" form:"csrf_token"`
	// FormError map[string]error
}


func (u User) Validate() error {
	vre := validation.Required.Error
	len := validation.Length(2, 20).Error
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name,
			vre("The name  is required"),
			len("The name field must be between 2 to 20 characters."),
		),
		validation.Field(&u.Email, vre("The Email  is required")),
		validation.Field(&u.Password, vre("The Password  is required")),
	)
}
