package handler

import (
	"auth-service/model"
	"auth-service/service"
	"auth-service/utils"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	service *service.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON at login request", http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Userpassword, []byte(req.Userpassword)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("JWT generation failed: %v", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON at register request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Userpassword), 0)
	if err != nil {
		log.Printf("Hashed password generation failed: %v", err)
		http.Error(w, "Failed to process registration", http.StatusInternalServerError)
		return
	}

	user := model.User{
		Username:     req.Username,
		Userpassword: hashedPassword,
	}
	exist, err := h.service.Register(&user)
	if err != nil {
		log.Printf("Exist user check failed: %v", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	if exist {
		http.Error(w, "User with such username already exist", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
