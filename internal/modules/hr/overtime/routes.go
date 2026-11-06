package overtime

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreatePolicy)
	router.GET("", handler.FindAllPolicies)
	router.GET("/:id", handler.FindPolicyByID)
	router.PUT("/:id", handler.UpdatePolicy)
	router.DELETE("/:id", handler.DeletePolicy)
}
