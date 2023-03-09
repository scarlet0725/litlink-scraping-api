package middleware

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
)

type Auth interface {
	Middleware() gin.HandlerFunc
}

type auth struct {
	user repository.User
}

func NewAuthMiddleware(user repository.User) Auth {
	return &auth{
		user: user,
	}
}

func (a *auth) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("X-API-KEY")
		if key == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"ok":      false,
				"message": "API Key is required",
			})
			return
		}

		sha512 := sha512.Sum512([]byte(key))
		k := hex.EncodeToString(sha512[:])

		user, err := a.user.GetUserByAPIKey(ctx, string(k))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"ok":      false,
				"message": "API Key is invalid",
			})
			return
		}

		if !user.IsAdminVerified {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"ok":      false,
				"message": "User is not verified",
			})
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
