package infrastructure

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/scarlet0725/prism-api/adapter"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/framework"
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
	router       *gin.Engine
	redis        *redis.Client
	ent          *ent.Client
	prismAPIHost string
}

func NewGinRouter(logger framework.Logger, ent *ent.Client, redis *redis.Client) (GinRouter, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(logger.GinLogger(), gin.Recovery())

	router := &ginRouter{
		router:       r,
		ent:          ent,
		redis:        redis,
		prismAPIHost: "",
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

	cache := NewRedisManager(r.redis)
	httpClient := NewHTTPClient()
	fetchController := controller.NewFetchController(httpClient, cache)

	docParser := parser.NewParser()
	serializer := selializer.NewResponseSerializer()

	userRepository := NewUserRepository(r.ent)
	artistRepository := NewArtistRepository(r.ent)
	eventRepository := NewEventRepository(r.ent)
	venueRepository := NewVenueRepository(r.ent)
	oAuthRepository := NewOAuthRepository(r.ent)

	googleOAuth := framework.NewGoogleOAuth(oauthConfig)

	eventUsecase := usecase.NewEventUsecase(eventRepository, artistRepository, venueRepository, fetchController, docParser, serializer, parser.NewJsonParser(), random)
	userUsecase := usecase.NewUserUsecase(userRepository, random, googleOAuth, NewGoogleCalenderClient)
	artistUsecase := usecase.NewArtistUsecase(artistRepository, random)
	venueUsecase := usecase.NewVenueUsecase(venueRepository, random)
	oauthUsecase := usecase.NewOAuthUsecase(oAuthRepository, random, googleOAuth)

	event := adapter.NewEventAdapter(eventUsecase)
	user := adapter.NewUserAdapter(userUsecase, eventUsecase)
	artist := adapter.NewArtistAdapter(artistUsecase)
	venue := adapter.NewVenueAdapter(venueUsecase)
	oauth := adapter.NewOAuthAdapter(oauthUsecase)

	v1 := r.router.Group("/v1")

	auth := middleware.NewAuthMiddleware(userRepository)

	userEndpoint := v1.Group("/user")
	eventEndpoint := v1.Group("/event")
	artistEndpoint := v1.Group("/artist")
	venueEndpoint := v1.Group("/venue")
	adminEndpoint := v1.Group("/admin")
	oauthEndpoint := v1.Group("/oauth")
	authEndpoint := v1.Group("/auth")
	searchEndpoint := v1.Group("/search")

	v1.POST("/register", user.Register)

	userEndpoint.Use(auth.Middleware())
	userEndpoint.GET("/", user.GetMe)
	userEndpoint.GET("/me", user.GetMe)
	userEndpoint.DELETE("/delete", user.Delete)
	userEndpoint.POST("/google", oauth.GoogleLinkage)
	userEndpoint.POST("/calendar", user.CreateExternalCalendar)
	userEndpoint.POST("/event", user.RegistrationEvent)

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

	authEndpoint.POST("/key", user.CreateAPIKey)

	searchEndpoint.Use(auth.Middleware())
	searchEndpoint.GET("/", event.SearchEvents)

	return nil
}

func (r *ginRouter) SetAPIHost(host string) {
	if host == "" {
		r.prismAPIHost = "http://localhost:8080"
		return
	}
	r.prismAPIHost = host

}
