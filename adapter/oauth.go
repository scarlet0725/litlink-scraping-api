package adapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/adapter/util"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

type OAuthAdapter interface {
	GoogleOAuthCallback(ctx *gin.Context)
	GoogleLinkage(ctx *gin.Context)
}

type oAuthAdapter struct {
	oauth usecase.OAuth
}

func NewOAuthAdapter(oauthUsecase usecase.OAuth) OAuthAdapter {
	return &oAuthAdapter{
		oauth: oauthUsecase,
	}
}

func (a *oAuthAdapter) GoogleOAuthCallback(ctx *gin.Context) {
	var req model.GoogleOauthCallback

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid request"})
		return
	}

	err := a.oauth.GoogleOAuthCallback(&req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "success"})

}

func (a *oAuthAdapter) GoogleLinkage(ctx *gin.Context) {
	var appErr *model.AppError

	user, err := util.GetUserFromContext(ctx)

	if err != nil {
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Invalid api Key"})
		return
	}

	resp, err := a.oauth.GoogleLinkage(user)

	if err != nil {
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "authenticate_url": resp.AuthURL, "message": "Please authenticate with google"})

}
