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
	MergeEvents(ctx *gin.Context)
	SearchEvents(ctx *gin.Context)
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
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
	var req model.UpdateEvent

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "invalid request"})
		return
	}

	eventID := ctx.Param("event_id")
	req.EventID = eventID

	_, err := a.event.UpdateEvent(&req)

	if err != nil {
		var appErr *model.AppError
		if errors.As(err, &appErr) {
			ctx.AbortWithStatusJSON(appErr.Code, gin.H{"ok": false, "error": appErr.Msg})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"updated": req,
		"message": "Successfully updated",
	})

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
	params := ctx.Param("event_id")

	event, err := a.event.GetEventByID(params)
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
	var req model.CrawlerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
	}

	result, err := a.event.CreateArtistEventsFromCrawlData(req.ArtistID)

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

func (a *eventAdapter) MergeEvents(ctx *gin.Context) {
	var req model.MergeEvent

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}

	merged, err := a.event.MergeEvents(&req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Merge Error",
		})
		return
	}

	ctx.JSON(200, gin.H{"ok": true, "events": merged})
}

func (a *eventAdapter) SearchEvents(ctx *gin.Context) {
	var req model.EventSearchQuery

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ok": false, "error": "param error"})
		return
	}

	events, err := a.event.SearchEvents(&req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"ok":    false,
			"error": "Search Error",
		})
		return
	}

	ctx.JSON(200, gin.H{"ok": true, "events": events})
}
