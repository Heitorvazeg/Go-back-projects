package user

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	r := Repository{db}
	return &r
}

func (r *Repository) CreateUsers(u *User) error {
	query := "INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, u.Nome, u.Email, u.Senha)
	return err
}

func (r *Repository) FindByEmail(u *User) (*User, error) {
	query := "SELECT id, nome, email, senha FROM usuarios WHERE email = ?"

	row := r.DB.QueryRow(query, u.Email)

	var user User

	err := row.Scan(&user.Id, &user.Nome, &user.Email, &user.Senha)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
