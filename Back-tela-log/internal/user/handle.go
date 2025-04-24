package user

import (
	"encoding/json"
	"net/http"
	"time"

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
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Erro ao ler JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	nome, email, senha := u.Nome, u.Email, u.Senha

	if err := val.Validate(nome, email, senha); err != nil {
		http.Error(w, "Erro: "+err.Error(), http.StatusBadRequest)
		return
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

	if err := h.Service.Repo.CreateUsers(&u); err != nil {
		http.Error(w, "Erro ao criar usuário"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário cadastrado com sucesso!"))

	NewLog(h, time.Now(), r.Method, r.URL, w.Header())
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Erro ao decodificar JSON! "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.Repo.FindByEmail(&u)

	if err != nil {
		http.Error(w, "Erro ao procurar email! "+err.Error(), http.StatusBadRequest)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Usuário não existe!"))
		return
	}

	if u.Email == user.Email && u.Senha == user.Senha {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Login realizado com sucesso!"))

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Email ou senha incorretas!"))

	}

	NewLog(h, time.Now(), r.Method, r.URL, w.Header())
}
