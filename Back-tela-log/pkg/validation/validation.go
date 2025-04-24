package val

import (
	"fmt"
	"regexp"
	"strings"
)

func Validate(nome, email, senha string) error {
	if err := validateRequiredFields(nome, email, senha); err != nil {
		return fmt.Errorf("campo(s) vazio(s)")
	}

	if !validateEmailFormat(email) {
		return fmt.Errorf("email inválido")
	}

	if !validatePassword(senha) {
		return fmt.Errorf("senha fraca")
	}

	return nil
}

func validateRequiredFields(nome, email, senha string) error {
	if nome == "" || email == "" || senha == "" {
		return fmt.Errorf("preencha os campos obrigatórios")
	}
	return nil
}

func validateEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(strings.TrimSpace(email))
}

func validatePassword(senha string) bool {
	return len(strings.TrimSpace(senha)) >= 7
}
