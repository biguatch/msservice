package msservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/biguatch/mslog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ServiceInterface interface {
	// Appends the given middlewares
	AddMiddleware(mws ...mux.MiddlewareFunc)
	// Adds the spefic route to router
	AddRouteHandler(path string, handler http.Handler, methods ...string)
	// Adds the spefic route to router
	AddRouteHandlerFunc(path string, f func(http.ResponseWriter, *http.Request), methods ...string)
	// Starts the listening
	ServeAPI(ctx context.Context)
	// Returns the logger
	GetLogger() *mslog.Logger
	// Creates the router
	crearteRouter() *mux.Router
}

type Service struct {
	// Name of the service for reference
	name string
	// Config options for this service
	config *Config
	// Router
	router *mux.Router
	// Logger
	logger *mslog.Logger
}

func NewService(name string, config *Config, logger *mslog.Logger) (service *Service, err error) {
	service = &Service{
		name:   name,
		config: config,
		logger: logger,
	}

	service.router = service.crearteRouter()

	return service, nil
}

func (service *Service) crearteRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	return router
}

func (service *Service) AddMiddleware(mws ...mux.MiddlewareFunc) {
	service.router.Use(mws...)
}

func (service *Service) AddRouteHandler(path string, handler http.Handler, methods ...string) {
	service.router.Handle(path, handler).Methods(methods...)
}

func (service *Service) AddRouteHandlerFunc(path string, f func(http.ResponseWriter, *http.Request), methods ...string) {
	service.router.HandleFunc(path, f).Methods(methods...)
}

func (service *Service) GetLogger() *mslog.Logger {
	return service.logger
}

func (service *Service) ServeAPI(ctx context.Context) {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", service.config.Port),
		Handler:     cors(service.router),
		ReadTimeout: 2 * time.Minute,
	}

	// The main go routine will wait until it receives
	// any message to this channel or until the channel is closed
	// This is how we are making the main go routine sync otherwise
	// main routine will exit and all sub routines (ListenAndServe)
	// will be terminates
	done := make(chan struct{})

	go func() {
		// This go routine will wait until the ctx channel is
		// closed. Once it is closed, it will shutdown the server
		// and will close the done channel, which in return will cause
		// main go rounte to exit
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			service.GetLogger().Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://0.0.0.0:%d", service.config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		service.GetLogger().Error(err)
	}

	// Either we received a message from this message (not our case here)
	// or the channel is closed (our case)
	<-done
}
