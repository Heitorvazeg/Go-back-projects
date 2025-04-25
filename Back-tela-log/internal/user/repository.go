package user

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	r := Repository{db}
	return &r
}

func (r *Repository) CreateUsers(u *User, senha string) error {
	query := "INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)"

	_, err := r.DB.Exec(query, u.Nome, u.Email, senha)
	return err
}

func (r *Repository) FindByEmail(u *User) (*User, error) {
	query := "SELECT nome, email, senha FROM usuarios WHERE email = ?"

	row := r.DB.QueryRow(query, u.Email)

	var user User

	err := row.Scan(&user.Nome, &user.Email, &user.Senha)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateLog(lg *Log) error {
	query := "INSERT INTO log (time, method, url, status, descricao) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, lg.Time, lg.Method, lg.Url, lg.Status, lg.Response)
	return err
}
