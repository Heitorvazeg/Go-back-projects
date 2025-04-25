package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	s := Service{repo}
	return &s
}

func (s *Service) EmailExists(u *User) (bool, error) {
	user, err := s.Repo.FindByEmail(u)

	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}

func (s *Service) SenhaCrypt(u *User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Erro ao criptografar senha!")
		return "", err
	}

	return string(hash), nil
}

func (s *Service) CorrectPassword(senhaCrypt, senhaDescrypt string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(senhaCrypt), []byte(senhaDescrypt))

	if err != nil {
		return false, err
	}
	return true, nil
}
