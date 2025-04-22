package val

import (
	"fmt"
	"regexp"
)

func Validate(email, senha, nome string) error {
	err := validateRequiredFields(nome, email, senha)

	if err != nil {
		return fmt.Errorf("os campos não podem estar vazios")
	}

	boolean := validateEmailFormat(email)

	if !boolean {
		return fmt.Errorf("email inválido")
	}

	boolean = validatePassword(senha)

	if !boolean {
		return fmt.Errorf("senha muito curta")
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
	return re.MatchString(email)
}

func validatePassword(senha string) bool {
	return len(senha) >= 7
}
