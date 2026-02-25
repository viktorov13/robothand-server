package auth

import (
	"encoding/json"
	"net/http"
	"robot-server/internal/models"
)

type Handler struct {
	Service *Service
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	token, uuid, err := h.Service.Register(req)
	if err != nil {
		http.Error(w, "Email exists", 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"uuid":  uuid,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&req)

	token, uuid, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"uuid":  uuid,
	})
}

func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "logged out"})
}
