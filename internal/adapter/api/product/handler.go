package product

import (
	"github.com/gin-gonic/gin"
	"github.com/phisher13/go-api/internal/domain/product/entity"
	"github.com/phisher13/go-api/internal/domain/product/service"
)

type handler struct {
	srv service.ProductService
}

func (h *handler) Update(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	var productDTO entity.ProductDTO
	if err := ctx.BindJSON(&productDTO); err != nil {
		ctx.JSON(400, err.Error())
	}
	err := h.srv.Update(uuid, productDTO)
	NewErrorHandler(err, ctx)

	ctx.JSON(200, productDTO)
}

func (h *handler) Delete(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	err := h.srv.Delete(uuid)
	NewErrorHandler(err, ctx)

	ctx.JSON(204, gin.H{"status": "deleted"})
}

func (h *handler) GetProducts(ctx *gin.Context) {
	userUUID, err := getUserUUID(ctx)
	NewErrorHandler(err, ctx)
	uuid := ctx.Query("uuid")
	if uuid != "" {
		product, err := h.srv.GetByUUID(uuid)
		NewErrorHandler(err, ctx)
		ctx.JSON(200, product)
	} else {
		products, err := h.srv.GetAll(userUUID)
		NewErrorHandler(err, ctx)
		ctx.JSON(200, products)
	}
}

func (h *handler) Create(ctx *gin.Context) {
	userUUID, err := getUserUUID(ctx)
	NewErrorHandler(err, ctx)

	var productDTO entity.ProductDTO
	if err := ctx.BindJSON(&productDTO); err != nil {
		ctx.JSON(400, err.Error())
	}
	uuid, err := h.srv.Create(productDTO, userUUID)

	NewErrorHandler(err, ctx)

	ctx.JSON(201, gin.H{"uuid": uuid})
}

func NewProductHandler(srv service.ProductService) ProductHandler {
	return &handler{
		srv: srv,
	}
}
