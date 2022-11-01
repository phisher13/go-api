package auth

import "github.com/gin-gonic/gin"

func (h *handler) InitRoutes(router *gin.Engine) {
	api := router.Group("api/")
	{
		api.POST("sign-in/", h.Login)
		api.POST("sign-up/", h.Register)
	}
}
