package middleware

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type Auth interface {
	Middleware() gin.HandlerFunc
}

type auth struct {
	user repository.User
	jwt  framework.JWTHandler
}

func NewAuthMiddleware(user repository.User, jwt framework.JWTHandler) Auth {
	return &auth{
		user: user,
		jwt:  jwt,
	}
}

func (a *auth) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *model.User
		var err error

		authHeader := ctx.GetHeader("Authorization")
		key := ctx.GetHeader("X-API-KEY")

		switch {
		case authHeader != "":
			token, err := a.jwt.ValidateToken(authHeader)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"ok":      false,
					"message": "Authorization header is invalid",
				})
				return
			}

	



			user, err = a.user.GetUser(ctx, token.UserID)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"ok":      false,
					"message": "Authorization header is invalid",
				})
				return
			}

		case key != "":
			sha512 := sha512.Sum512([]byte(key))
			k := hex.EncodeToString(sha512[:])

			user, err = a.user.GetUserByAPIKey(ctx, string(k))
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"ok":      false,
					"message": "API Key is invalid",
				})
				return
			}

			//
			fallthrough
		default:
			if user == nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"ok":      false,
					"message": "Authorization or X-API-KEY header is required",
				})
				return
			}

		}

		if !user.IsVerified() {
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
