package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationToken = "session"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	tokenString, _ := c.Cookie("session")
	if tokenString == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty token")
		return
	}

	userId, err := h.services.Authorization.ParseToken(tokenString)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// func getUserId(c *gin.Context) (string, error) {
// 	id, ok := c.Get(userCtx)
// 	if !ok {
// 		return "", errors.New("user id not found")
// 	}

// 	idInt, ok := id.(string)
// 	if !ok {
// 		return "", errors.New("user id is of invalid type")
// 	}

// 	return idInt, nil
// }
