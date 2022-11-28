package adapter

import (
	"errors"
	"net/http"
	"net/mail"

	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

var (
	usernameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_]{3,256}$`)
)

type UserAdapter interface {
	Register(ctx *gin.Context)
	//Login(ctx *gin.Context)
	//GetUser(ctx *gin.Context)
	//CreateUser(ctx *gin.Context)
	//UpdateUser(ctx *gin.Context)
	Delete(ctx *gin.Context)
	//CreateAPIKey(ctx *gin.Context)
	GetMe(ctx *gin.Context)
}

type userAdapter struct {
	user usecase.User
}

func NewUserAdapter(user usecase.User) UserAdapter {
	return &userAdapter{
		user: user,
	}
}

func (a *userAdapter) Register(ctx *gin.Context) {
	var req model.RegisterUser

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid request"})
		return
	}

	//ユーザーネームが半角英数字とアンダーバーのみかどうか
	if !usernameRegexp.MatchString(req.Username) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Usernames must be between 3 and 256 alphanumeric characters or underscored."})
		return
	}

	//メールアドレスが正しいかどうか
	_, err := mail.ParseAddress(req.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid email address"})
		return
	}

	//パスワードがASCII文字のみであるかチェック
	if !(utf8.ValidString(req.Password) && utf8.RuneCountInString(req.Password) == len(req.Password)) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "password must be ASCII characters"})
		return
	}

	//パスワードが8文字以上64文字以下であるかチェック
	if len(req.Password) < 8 || len(req.Password) > 64 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Password must be between 8 and 64 characters"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "failed to hash password"})
		return
	}

	user := &model.User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        hash,
		IsAdminVerified: false,
	}

	result, err := a.user.CreateUser(user)

	if err != nil {
		var appErr *model.AppError
		if ok := errors.As(err, &appErr); ok {
			ctx.JSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "user": result})

}

func (c *userAdapter) GetMe(ctx *gin.Context) {
	k := ctx.GetHeader("X-Api-Key")
	user, err := c.user.GetUserByAPIKey(k)
	if err != nil {
		var appErr *model.AppError
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid api key"})
		return
	}

	resp := model.UserResponse{
		OK:   true,
		User: user,
	}
	ctx.JSON(200, resp)
}

func (c *userAdapter) Delete(ctx *gin.Context) {
	k := ctx.GetHeader("X-Api-Key")
	user, err := c.user.GetUserByAPIKey(k)

	if err != nil {
		var appErr *model.AppError
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"ok": false, "error": "User not found"})
		return
	}

	if user.DeleteProtected {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"ok": false, "error": "User can not be deleted"})
		return
	}

	err = c.user.DeleteUser(user)

	if err != nil {
		var appErr *model.AppError
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "InternalServerError"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "User deleted"})

}
