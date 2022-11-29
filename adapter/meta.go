package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetaController struct {
}

func NewMetaController() MetaController {
	return MetaController{}
}

func (c *MetaController) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "ok"})
}

func (c *MetaController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "ok",
	})
}

func (c *MetaController) NoMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, gin.H{
		"ok":    false,
		"error": "Method Not Allowed",
	})
}

func (c *MetaController) NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"ok":    false,
		"error": "Not Found",
	})
}
