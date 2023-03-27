package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

type Venue interface {
	CreateArtist(*gin.Context)
}

type VenueAdapter struct {
	venueUsecase usecase.Venue
}

func NewVenueAdapter(venueUsecase usecase.Venue) *VenueAdapter {
	return &VenueAdapter{
		venueUsecase: venueUsecase,
	}
}

func (a *VenueAdapter) CreateVenue(ctx *gin.Context) {
	var req model.CreateVenue
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"ok":      false,
			"message": "Invalid request",
		})
		return
	}

	venue := &model.Venue{
		Name:        req.Name,
		Description: req.Description,
		WebSite:     req.WebSite,
		Postcode:    req.Postcode,
		Prefecture:  req.Prefecture,
		City:        req.City,
		Street:      req.Street,
	}

	result, err := a.venueUsecase.CreateVenue(ctx, venue)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"ok":      false,
			"message": "Failed to create venue",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":      true,
		"message": "Success",
		"venue":   result,
	})
}
