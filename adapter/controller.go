package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

type EventController struct {
	event usecase.EventApplication
}

func NewController(event usecase.EventApplication) EventController {
	return EventController{
		event: event,
	}
}

func (c *EventController) GetEventsByArtistName(ctx *gin.Context) {
	name := ctx.Param("arist_name")

	events, err := c.event.GetEventsByArtistName(name)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	res := model.EventAPIResponse{
		OK:   true,
		Data: events,
	}

	ctx.JSON(200, res)

}
