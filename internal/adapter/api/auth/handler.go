package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/phisher13/go-api/internal/domain/auth/entity"
	"github.com/phisher13/go-api/internal/domain/auth/service"
	"net/http"
)

type handler struct {
	srv service.AuthService
}

func NewAuthorizationHandler(srv service.AuthService) AuthorizationHandler {
	return &handler{srv: srv}
}

func (h *handler) Login(ctx *gin.Context) {
	var dto entity.UserDTO
	if err := ctx.BindJSON(&dto); err != nil {
		NewErrorHandler(err, ctx)
	}
	token, err := h.srv.GenerateToken(dto.Username, dto.Password)
	NewErrorHandler(err, ctx)

	ctx.JSON(http.StatusOK, gin.H{"Authorization": token})
}

func (h *handler) Register(ctx *gin.Context) {
	var dto entity.UserDTO
	if err := ctx.BindJSON(&dto); err != nil {
		NewErrorHandler(err, ctx)
	}
	userUuid, err := h.srv.CreateUser(dto)
	NewErrorHandler(err, ctx)

	ctx.JSON(http.StatusCreated, gin.H{"user_uuid": userUuid})
}
