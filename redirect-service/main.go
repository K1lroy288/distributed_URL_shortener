package main

import (
	"fmt"
	"net/http"
	"redirect-service/client"
	"redirect-service/config"
	"redirect-service/handler"
	"redirect-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	redisAddr := cfg.RedisHost + ":" + cfg.RedisPort
	shortenerAddr := "http://" + cfg.ShortenerHost + ":" + cfg.ShortenerPort

	redis := client.NewRedisClient(redisAddr)
	shortener := client.NewShortenerClient(shortenerAddr)
	service := service.NewRedirectService(redis, shortener)
	handler := handler.NewRedirectHandler(service)

	r := gin.Default()

	r.GET("/redirect/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Redirect service is up!")
	})

	r.GET("/:shortCode", handler.Resolve)

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
