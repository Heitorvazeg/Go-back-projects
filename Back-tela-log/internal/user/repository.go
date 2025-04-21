package user

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateUsers(u *User) error {
	_, err := r.DB.Exec("INSERT INTO usuarios (nome, email) VALUES")
	return err
}

func (r *Repository) FindByEmail(u *User) (*User, error) {

}
