package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/scarlet0725/prism-api/adapter"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/middleware"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/repository"
	"github.com/scarlet0725/prism-api/selializer"
	"github.com/scarlet0725/prism-api/usecase"
)

type GinRouter interface {
	Serve(addr string) error
	SetRoute()
}

type ginRouter struct {
	fetch      controller.FetchController
	paser      parser.DocParser
	selializer selializer.ResponseSerializer
	router     *gin.Engine
	db         repository.DB
}

func NewGinRouter(fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer, db repository.DB) GinRouter {
	r := gin.Default()

	router := &ginRouter{
		fetch:      fetch,
		paser:      parser,
		selializer: selializer,
		router:     r,
		db:         db,
	}
	router.SetMeta()
	router.SetRoute()
	r.HandleMethodNotAllowed = true

	return router

}

func (r *ginRouter) Serve(addr string) error {
	return r.router.Run(addr)
}

func (r *ginRouter) SetMeta() {
	m := adapter.NewMetaController()
	r.router.GET("/health", m.HealthCheck)
	r.router.GET("/healthz", m.HealthCheck)

	r.router.NoMethod(m.NoMethod)
	r.router.NoRoute(m.NoRoute)

}

func (r *ginRouter) SetRoute() {
	eventUsecase := usecase.NewEventApplication(r.db, r.fetch, r.paser, r.selializer, parser.NewJsonParser())
	userUsecase := usecase.NewUserApplication(r.db)
	artistUsecase := usecase.NewArtistUsecase(r.db)
	venueUsecase := usecase.NewVenueUsecase(r.db)

	event := adapter.NewEventAdapter(eventUsecase)
	user := adapter.NewUserAdapter(userUsecase)
	artist := adapter.NewArtistAdapter(artistUsecase)
	venue := adapter.NewVenueAdapter(venueUsecase)
	v1 := r.router.Group("/v1")

	auth := middleware.NewAuthMiddleware(r.db)

	userEndpoint := v1.Group("/user")
	eventEndpoint := v1.Group("/event")
	artistEndpoint := v1.Group("/artist")
	venueEndpoint := v1.Group("/venue")
	adminEndpoint := v1.Group("/admin")

	v1.POST("/register", user.Register)

	userEndpoint.Use(auth.Middleware())
	userEndpoint.GET("/", user.GetMe)
	userEndpoint.GET("/me", user.GetMe)
	userEndpoint.DELETE("/delete", user.Delete)

	//v1.GET("events/:arist_name", event.GetEventsByArtistName)
	eventEndpoint.GET("/:event_id", event.GetEventByID)
	eventEndpoint.DELETE("/:event_id", event.DeleteEvent)
	eventEndpoint.POST("/:event_id", event.UpdateEvent)
	eventEndpoint.POST("/", event.CreateEvent)

	artistEndpoint.POST("/", artist.CreateArtist)
	artistEndpoint.POST("/events/auto_update", event.CreateArtistEventsFromCrawlData)

	venueEndpoint.POST("/", venue.CreateVenue)

	adminEndpoint.POST("/verify_account", user.Verify)

}
