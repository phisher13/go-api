package product

import "github.com/gin-gonic/gin"

func (h *handler) InitRoutes(router *gin.Engine) {
	api := router.Group("api/")
	{
		product := api.Group("product/", h.userIdentity)
		{
			product.GET("/", h.GetProducts)
			product.POST("new/", h.Create)
			product.DELETE("/", h.Delete)
			product.PUT("/", h.Update)
		}
	}
}
