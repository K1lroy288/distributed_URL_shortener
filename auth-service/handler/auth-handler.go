package handler

import (
	"auth-service/model"
	"auth-service/service"
	"auth-service/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at login request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := h.service.Login(req.Username)
	if err != nil {
		log.Printf("Invalid username or password: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Userpassword, []byte(req.Password)); err != nil {
		log.Printf("Invalid username or password: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("JWT generation failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	response := map[string]string{"token": token}
	ctx.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at register request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON at register request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hashed password generation failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
		return
	}

	user := model.User{
		Username:     req.Username,
		Userpassword: hashedPassword,
	}
	exist, err := h.service.Register(&user)
	if err != nil {
		log.Printf("Exist user check failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	if exist {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User with such username already exists"})
		return
	}

	ctx.Status(http.StatusCreated)
}
