package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("/login", handler.Login)
	router.POST("/refresh", handler.RefreshToken)
	router.POST("/logout", handler.Logout)
}
