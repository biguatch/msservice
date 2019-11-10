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
	AddMiddleware(mws ...mux.MiddlewareFunc)
	AddRouteHandler(path string, handler http.Handler, methods ...string)
	AddRouteHandlerFunc(path string, f func(http.ResponseWriter, *http.Request), methods ...string)
	ServeAPI(ctx context.Context)
	GetLogger() *mslog.Logger
}

type Service struct {
	name   string
	config *Config
	router *mux.Router
	logger *mslog.Logger
}

func NewService(name string, config *Config, logger *mslog.Logger) (service *Service, err error) {
	service = &Service{
		name:   name,
		config: config,
		logger: logger,
	}

	service.router = crearteRouter(service)

	return service, nil
}

func crearteRouter(service *Service) *mux.Router {
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

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://0.0.0.0:%d", service.config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}
