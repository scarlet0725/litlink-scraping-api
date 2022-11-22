package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

type ArtistAdapter interface {
	CreateArtist(*gin.Context)
}

type artistAdapter struct {
	artist usecase.ArtistUsecase
}

func NewArtistAdapter(artist usecase.ArtistUsecase) ArtistAdapter {
	return &artistAdapter{artist: artist}
}

func (a *artistAdapter) CreateArtist(ctx *gin.Context) {
	var req model.CreateArtist
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "msg": "Bad Request"})
		return
	}

	if req.Name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "msg": "Artist name is required"})
		return
	}

	artist := &model.Artist{
		Name: req.Name,
		URL:  req.URL,
	}

	result, err := a.artist.CreateArtist(artist)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "msg": "Failed to create artist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "artist": result})

}
