package util

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
)

func GetUserFromContext(ctx *gin.Context) (*model.User, error) {
	val, ok := ctx.Get("user")
	if !ok {
		return nil, errors.New("user not found")
	}

	user, ok := val.(*model.User)
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
