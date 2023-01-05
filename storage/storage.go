package storage

import (
	_ "github.com/golang/mock/mockgen/model"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	CreateUser(u *User) (*User, error)
	GetUser(id int64) (*User, error)
}

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type storagePg struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) StorageI {
	return &storagePg{
		db: db,
	}
}

func (ur *storagePg) CreateUser(user *User) (*User, error) {
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			phone_number,
			email
		) VALUES($1, $2, $3, $4)
		RETURNING id
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
	)

	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *storagePg) GetUser(id int64) (*User, error) {
	var result User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email
		FROM users
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
