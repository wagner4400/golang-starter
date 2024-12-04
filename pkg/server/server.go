package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"lawise-go/internal/entrypoint/http/router"
	"log"
	"net/http"
)

type HttpApi interface {
	Method(method, path string, handler gin.HandlerFunc)
	Shutdown(ctx context.Context) error
	Run(ctx context.Context)
	ServeHTTP(rw http.ResponseWriter, request *http.Request)
	GetRouter() *router.Router
}

type serverApi struct {
	config *Config
	server *http.Server
	engine *gin.Engine
	router *router.Router
}

func NewHttpServerApi(config *Config) HttpApi {
	engine := gin.New()

	server := &http.Server{
		Addr:         config.Address,
		Handler:      engine,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	ServerRouter := router.NewRouter(engine)

	srv := serverApi{config, server, engine, ServerRouter}
	//St a lower memory limit for multipart forms (default is 32 MiB)
	srv.engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	return &srv
}

func (s *serverApi) Method(method, path string, handler gin.HandlerFunc) {
	s.engine.Handle(method, path, handler)
}

func (s *serverApi) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.config.ShutdownTimeout)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		ctx.Done()
		log.Println("starting shutdown http server")
		return err
	}
	log.Println("http server stopped")
	return nil
}

func (s *serverApi) Run(ctx context.Context) {
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err.Error())
		ctx.Done()
	}
}

func (s *serverApi) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	s.engine.ServeHTTP(rw, request)
}

func (s *serverApi) GetRouter() *router.Router {
	return s.router
}
