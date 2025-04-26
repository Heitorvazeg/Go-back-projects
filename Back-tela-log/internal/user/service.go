package user

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

func (s *Service) newToken(u *User) (string, error) {
	id, err := s.Repo.findID(u)

	if id < 0 {
		if err != nil {
			return "", err
		}
		return "", nil
	}

	if err = godotenv.Load("C:/Users/heito/OneDrive/Documentos/Programas/Go-back-projects/Back-tela-log/config/login.env"); err != nil {
		return "", err
	}

	jwtkey := os.Getenv("JWT_KEY")

	claim := jwt.MapClaims{
		"user_id": id,
		"exp":     int(time.Now().Add(time.Hour * 24).Unix()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(jwtkey))
}
