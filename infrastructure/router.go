package infrastructure

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/scarlet0725/prism-api/adapter"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/middleware"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/selializer"
	"github.com/scarlet0725/prism-api/usecase"
)

type GinRouter interface {
	Serve(addr string) error
	SetRoute() error
}

type ginRouter struct {
	fetch        controller.FetchController
	paser        parser.DocParser
	selializer   selializer.ResponseSerializer
	router       *gin.Engine
	db           repository.DB
	prismAPIHost string
}

func NewGinRouter(logger *zap.Logger, fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer, db repository.DB) (GinRouter, error) {
	r := gin.New()

	r.Use(middleware.Logger(logger), gin.Recovery())

	router := &ginRouter{
		fetch:      fetch,
		paser:      parser,
		selializer: selializer,
		router:     r,
		db:         db,
	}
	r.HandleMethodNotAllowed = true
	r.TrustedPlatform = gin.PlatformCloudflare

	router.SetAPIHost(os.Getenv("PRISM_API_HOST"))
	router.SetMeta()
	err := router.SetRoute()

	return router, err

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

func (r *ginRouter) SetRoute() error {
	oauthConfig, err := cmd.GetGoogleOAuthConfig(r.prismAPIHost)
	random := framework.NewRamdomIDGenerator()

	if err != nil {
		return err
	}

	eventUsecase := usecase.NewEventApplication(r.db, r.fetch, r.paser, r.selializer, parser.NewJsonParser(), random)
	userUsecase := usecase.NewUserApplication(r.db, random)
	artistUsecase := usecase.NewArtistUsecase(r.db, random)
	venueUsecase := usecase.NewVenueUsecase(r.db, random)
	oauthUsecase := usecase.NewOAuthApplication(r.db, random, framework.NewGoogleOAuth(oauthConfig))

	event := adapter.NewEventAdapter(eventUsecase)
	user := adapter.NewUserAdapter(userUsecase)
	artist := adapter.NewArtistAdapter(artistUsecase)
	venue := adapter.NewVenueAdapter(venueUsecase)
	oauth := adapter.NewOAuthAdapter(oauthUsecase)

	v1 := r.router.Group("/v1")

	auth := middleware.NewAuthMiddleware(r.db)

	userEndpoint := v1.Group("/user")
	eventEndpoint := v1.Group("/event")
	artistEndpoint := v1.Group("/artist")
	venueEndpoint := v1.Group("/venue")
	adminEndpoint := v1.Group("/admin")
	oauthEndpoint := v1.Group("/oauth")

	v1.POST("/register", user.Register)

	userEndpoint.Use(auth.Middleware())
	userEndpoint.GET("/", user.GetMe)
	userEndpoint.GET("/me", user.GetMe)
	userEndpoint.DELETE("/delete", user.Delete)
	userEndpoint.POST("/google", oauth.GoogleLinkage)

	//v1.GET("events/:arist_name", event.GetEventsByArtistName)
	eventEndpoint.Use(auth.Middleware())
	eventEndpoint.GET("/:event_id", event.GetEventByID)
	eventEndpoint.DELETE("/:event_id", event.DeleteEvent)
	eventEndpoint.POST("/:event_id", event.UpdateEvent)
	eventEndpoint.POST("/", event.CreateEvent)
	eventEndpoint.POST("/merge", event.MergeEvents)

	artistEndpoint.Use(auth.Middleware())
	artistEndpoint.POST("/", artist.CreateArtist)
	artistEndpoint.POST("/events/auto_update", event.CreateArtistEventsFromCrawlData)

	adminEndpoint.Use(auth.Middleware())
	venueEndpoint.POST("/", venue.CreateVenue)

	adminEndpoint.Use(auth.Middleware())
	adminEndpoint.POST("/verify_account", user.Verify)

	oauthEndpoint.GET("/google/callback", oauth.GoogleOAuthCallback)

	return nil
}

func (r *ginRouter) SetAPIHost(host string) {
	if host == "" {
		r.prismAPIHost = "http://localhost:8080"
		return
	}
	r.prismAPIHost = host

}
