package handler

import (
	"log"
	"net/http"
	"shortener-service/config"
	"shortener-service/model"
	"shortener-service/service"
	"shortener-service/utils"

	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	service *service.ShortenerService
}

func NewShortenerHandler(service *service.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{service: service}
}

func (h *ShortenerHandler) SaveCode(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	var req string
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at save code request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for {
		code, err := utils.GenerateShortCode()
		if err != nil {
			log.Printf("Generate short code error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		url := model.Url{
			Long_url:   req,
			Short_code: code,
			Owner_id:   claims["user_id"].(int),
		}

		exist, err := h.service.SaveCode(url)
		if err != nil {
			log.Printf("Save short code error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}

		if !exist {
			cfg := config.GetConfig()

			url := "http://" + cfg.Host + ":" + cfg.Port + "/short/" + code

			ctx.JSON(http.StatusCreated, gin.H{"url": url})
			break
		}
	}
}

func (h *ShortenerHandler) GetLink(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	_, err := utils.ValidateJWT(token)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	code := ctx.Param("shortCode")
	url, err := h.service.GetLink(code)
	if err != nil {
		log.Printf("error git link: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Such link don't exist"})
		return
	}

	if url == nil {
		log.Printf("error git link: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Such link don't exist"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.Long_url)
}
