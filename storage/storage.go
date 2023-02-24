package storage

import (
	"database/sql"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StudentFilter struct {
	SearchTerm string
}

type Student struct {
	ID        int          `db:"id" form:"-"`
	FirstName string       `db:"first_name" form:"first_name"`
	LastName  string       `db:"last_name" form:"last_name"`
	Class     int          `db:"class" form:"class"`
	Roll      int          `db:"roll" form:"roll"`
	Email     string       `db:"email" form:"email"`
	Password  string       `db:"password" form:"password"`
	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
	// CSRFToken string    `db:"-" form:"csrf_token"`
	// FormError map[string]error
}

func (s Student) Validate() error {
	vre := validation.Required.Error
	
	len := validation.Length(2, 20).Error
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName,
			vre("The FirstName  is required"),
			len("The FirstName field must be between 2 to 20 characters."),
		),
		validation.Field(&s.LastName,
			vre("The LastName  is required"),
			len("The LastName field must be between 2 to 20 characters."),
		),
		validation.Field(&s.Class, vre("The Class  is required")),
		validation.Field(&s.Roll, vre("The Roll  is required"),
		validation.Required, validation.Min(1), validation.Max(100).Error(" Roll must Positive  (1 to 100) "),),
		validation.Field(&s.Email, vre("The Email  is required")),
		validation.Field(&s.Password, vre("The Password  is required")),
	)
}

type UserFilter struct {
	SearchTerm string
}

type User struct {
	ID        int          `db:"id" form:"-"`
	Name      string       `db:"name" form:"name"`
	FirstName string       `db:"first_name" form:"first_name"`
	LastName  string       `db:"last_name" form:"last_name"`
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
	nameRule := validation.Match(
		regexp.MustCompile(`^[^\s]+$`)).
		Error("Name must not contain spaces")
	len := validation.Length(2, 20).Error
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			vre("The FirstName  is required"),
			len("The FirstName field must be between 2 to 20 characters."),
		),
		validation.Field(&u.LastName,
			vre("The LastName  is required"),
			len("The LastName field must be between 2 to 20 characters."),
		),
		validation.Field(&u.Name,
			vre("The name  is required"),
			len("The name field must be between 2 to 20 characters."), nameRule,
		),
		validation.Field(&u.Email, vre("The Email  is required")),
		validation.Field(&u.Password, vre("The Password  is required")),
	)
}

// Class

type Class struct {
	ID        int          `db:"id" form:"-"`
	ClassName string       `db:"class_name" form:"class_name"`
	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}

func (c Class) Validate() error {
	vre := validation.Required.Error
	classRule := validation.Match(
		regexp.MustCompile(`^Class [1-9]$|^Class 10$`)).
		Error("Class must be in the format 'Class [1-10]'")
	return validation.ValidateStruct(&c,
		validation.Field(&c.ClassName, vre("The Class Name  is required"),
			classRule),
	)
}

// Subject

type SubjectFilter struct {
	SearchTerm string
}

type Subject struct {
	ID       int    `db:"id" form:"-"`
	Class    int    `db:"class" form:"class"`
	Subject1 string `db:"subject1" form:"subject1"`

	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}

func (s Subject) Validate() error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&s,
		validation.Field(&s.Class,
			vre("The Class  is required"),
		),
		validation.Field(&s.Subject1, vre("The LastName  is required")),
	)
}

type StudentSubject struct {
	ID        int `db:"id" form:"-"`
	StudentID int `db:"student_id" form:"student_id"`
	SubjectID int `db:"subject_id" form:"subject_id"`
	Marks     int `db:"marks"`
	Mark      map[int]int
	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}

type MarkInputStore struct {
	ID        int          `db:"id" form:"id"`
	FirstName string       `db:"first_name" form:"first_name"`
	LastName  string       `db:"last_name" form:"last_name"`
	Class     int          `db:"class" form:"class"`
	Roll      int          `db:"roll" form:"roll"`
	SubjectID int          `db:"subject_id" form:"subject_id"`
	Subject1  string       `db:"subject1" form:"subject1"`
	CreatedAt time.Time    `db:"created_at" form:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" form:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" form:"deleted_at"`
}

type Sts struct {
	Mark map[int]int
}

type Result struct {
	ID        int    `db:"id" form:"id"`
	FirstName string `db:"first_name,omitempty" form:"first_name"`
	LastName  string `db:"last_name,omitempty" form:"last_name"`
	Class     int    `db:"class,omitempty" form:"class"`
	Roll      int    `db:"roll,omitempty" form:"roll"`
	Subject1  string `db:"subject1,omitempty" form:"subject1"`
	Marks     int    `db:"marks,omitempty"`
}

type AllResult struct {
	ID        sql.NullInt64  `db:"id" form:"id"`
	FirstName sql.NullString `db:"first_name" form:"first_name"`
	LastName  sql.NullString `db:"last_name" form:"last_name"`
	Class     sql.NullInt64  `db:"class" form:"class"`
	Roll      sql.NullInt64  `db:"roll" form:"roll"`
	Subject1  sql.NullString `db:"subject1" form:"subject1"`
	Marks     sql.NullInt64  `db:"marks"`
}


type MarkEdit struct {
	ID        int    `form:"ID" db:"id"`
	Marks     string `form:"Marks" db:"marks"`
}