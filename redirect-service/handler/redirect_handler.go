package handler

import (
	"log"
	"net/http"
	"redirect-service/service"

	"github.com/gin-gonic/gin"
)

type RedirectHandler struct {
	service *service.RedirectService
}

func NewRedirectHandler(service *service.RedirectService) *RedirectHandler {
	return &RedirectHandler{service: service}
}

func (h *RedirectHandler) Resolve(ctx *gin.Context) {
	code := ctx.Param("shortCode")
	url, err := h.service.Resolve(ctx, code)
	if err != nil {
		log.Printf("error git link: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}
