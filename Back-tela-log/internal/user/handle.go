package user

import (
	"encoding/json"
	"net/http"

	val "github.com/Heitorvazeg/Go-back-projects/Back-tela-log/pkg/validation"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	h := Handler{service}
	return &h
}

func (h *Handler) HandleCadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Erro ao ler JSON"+err.Error(), http.StatusBadRequest)
		return
	}
	nome, email, senha := u.Nome, u.Email, u.Senha

	err = val.Validate(nome, email, senha)

	if err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
	}

	yes, err := h.Service.EmailExists(&u)

	if err != nil {
		http.Error(w, "Erro ao validar email!", http.StatusBadRequest)
		return
	}

	if yes {
		http.Error(w, "Email já cadastrado!", http.StatusConflict)
		return
	}

	err = h.Service.Repo.CreateUsers(&u)

	if err != nil {
		http.Error(w, "Erro ao criar usuário"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário cadastrado com sucesso!"))
}
