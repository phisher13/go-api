package product

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/phisher13/go-api/internal/domain/auth/service"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, "token is empty")
		return
	}
	srv := service.Service{}
	userId, err := srv.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserUUID(c *gin.Context) (string, error) {
	uuid, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return uuidStr, nil
}
