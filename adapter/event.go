package adapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/selializer"
	"github.com/scarlet0725/prism-api/usecase"
)

const (
	EventDateLayout = "2006-01-02"
)

type EventAdapter interface {
	GetEvent(ctx *gin.Context)
	CreateEvent(ctx *gin.Context)
	UpdateEvent(ctx *gin.Context)
	DeleteEvent(ctx *gin.Context)
	GetEventsByArtistName(ctx *gin.Context)
	GetEventByID(ctx *gin.Context)
	CreateArtistEventsFromCrawlData(ctx *gin.Context)
}

type eventAdapter struct {
	event      usecase.Event
	selializer selializer.ResponseSerializer
}

func NewEventAdapter(eventController usecase.Event) EventAdapter {
	selializer := selializer.NewResponseSerializer()
	return &eventAdapter{
		selializer: selializer,
		event:      eventController,
	}
}

func (a *eventAdapter) GetEvent(ctx *gin.Context) {
}

func (a *eventAdapter) CreateEvent(ctx *gin.Context) {
	var req model.CreateEvent

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid request"})
		return
	}

	result, err := a.event.CreateEvent(&req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "event": result})

}

func (a *eventAdapter) UpdateEvent(ctx *gin.Context) {
	//TODO: 実装する
}

func (a *eventAdapter) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("event_id")

	event, err := a.event.GetEventByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Event not found",
		})
		return
	}

	if err := a.event.DeleteEvent(event); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "Successfully deleted"})

}

func (a *eventAdapter) GetEventsByArtistName(ctx *gin.Context) {
	name := ctx.Param("arist_name")

	events, err := a.event.GetEventsByArtistName(name)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, gin.H{"ok": true, "events": events})

}

func (a *eventAdapter) GetEventByID(ctx *gin.Context) {
	var params model.GetEvent

	if err := ctx.ShouldBind(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok": false, "error": "Parameter event_id is invalid",
		})
		return
	}

	event, err := a.event.GetEventByID(params.EventID)
	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"ok":    false,
			"error": "event not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"event": event,
	})
}

func (a *eventAdapter) CreateArtistEventsFromCrawlData(ctx *gin.Context) {
	artistID := ctx.Param("artist_id")

	result, err := a.event.CreateArtistEventsFromCrawlData(artistID)

	if err != nil {
		var appErr *model.AppError
		if ok := errors.As(err, &appErr); ok {
			ctx.AbortWithStatusJSON(appErr.Code, appErr.Msg)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":    false,
			"error": "Internal server error",
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Successfully created",
		"data":    result,
	})
}
