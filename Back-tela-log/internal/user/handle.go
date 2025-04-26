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

	ok, err := h.Service.EmailExists(&u)

	if err != nil {
		http.Error(w, "Erro ao validar email!", http.StatusBadRequest)
		return
	}

	if ok {
		http.Error(w, "Email já cadastrado!", http.StatusConflict)
		return
	}

	senhaSec, err := h.Service.SenhaCrypt(&u)

	if err != nil {
		http.Error(w, "Erro ao criptografar senha! "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.Repo.CreateUsers(&u, senhaSec); err != nil {
		http.Error(w, "Erro ao criar usuário"+err.Error(), http.StatusInternalServerError)
		return
	}

	lrw := NewLGRW(w)

	lrw.WriteHeader(http.StatusAccepted)
	lrw.Write([]byte("Email cadastrado com sucesso!"))

	NewLog(h, lrw, time.Now(), r.Method, r.URL)
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

	lrw := NewLGRW(w)

	if user == nil {
		lrw.WriteHeader(http.StatusConflict)
		lrw.Write([]byte("Usuário não existe!"))
		NewLog(h, lrw, time.Now(), r.Method, r.URL)
		return
	}

	ok, err := h.Service.CorrectPassword(user.Senha, u.Senha)

	if err != nil {
		http.Error(lrw, "Senha incorreta!", http.StatusUnauthorized)
		NewLog(h, lrw, time.Now(), r.Method, r.URL)
		return
	}

	if u.Email == user.Email && ok {
		token, err := h.Service.newToken(&u)

		if err != nil {
			http.Error(w, "Erro ao gerar token! "+err.Error(), http.StatusBadRequest)
		}

		if err := json.NewEncoder(lrw).Encode(map[string]string{
			"token":    token,
			"mensagem": "Login realizado com sucesso!",
		}); err != nil {
			http.Error(w, "Erro ao codificar JSON token! "+err.Error(), http.StatusBadRequest)
		}
	}

	NewLog(h, lrw, time.Now(), r.Method, r.URL)
}
