package auth

import (
	"database/sql"
	"robot-server/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateUser(u models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users VALUES (?, ?, ?, ?, ?)",
		u.UUID, u.Email, u.Password, u.Name, u.Surname,
	)
	return err
}

func (r *Repository) GetByEmail(email string) (*models.User, error) {
	row := r.DB.QueryRow(
		"SELECT uuid, email, password, name, surname FROM users WHERE email = ?",
		email,
	)

	var u models.User
	err := row.Scan(&u.UUID, &u.Email, &u.Password, &u.Name, &u.Surname)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
