package router

import (
	"github.com/gin-gonic/gin"

	"github.com/scarlet0725/prism-api/adapter"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/parser"
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
}

func NewGinRouter(fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer) GinRouter {
	r := gin.Default()

	router := &ginRouter{
		fetch:      fetch,
		paser:      parser,
		selializer: selializer,
		router:     r,
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
	event := usecase.NewEventApplication(gateway.NewDevDB(), r.fetch, r.paser, r.selializer, parser.NewJsonParser())
	controller := adapter.NewController(event)
	v1 := r.router.Group("/v1")

	v1.GET("events/:arist_name", controller.GetEventsByArtistName)
}
