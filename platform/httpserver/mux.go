package httpserver

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/newbie007fx/cinemas/platform/configuration"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HttpService struct {
	cs     *configuration.ConfigService
	route  *mux.Router
	server *http.Server
}

func New(cs *configuration.ConfigService) *HttpService {
	return &HttpService{
		cs: cs,
	}
}

func (rs *HttpService) Setup() error {
	conf := rs.cs.GetConfig()

	rs.route = mux.NewRouter()

	allowedHeaders := handlers.AllowedHeaders(conf.CorsConfig.AllowedHeaders)
	allowedOrigins := handlers.AllowedOrigins(conf.CorsConfig.AllowedOrigins)
	allowedMethods := handlers.AllowedMethods(conf.CorsConfig.AllowedMethods)

	handler := handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods, handlers.MaxAge(1000))(rs.route)
	if conf.App.LogLevel == "DEBUG" {
		handler = handlers.LoggingHandler(os.Stdout, handler)
	}

	rs.server = &http.Server{
		Handler:           handler,
		Addr:              fmt.Sprintf(":%d", conf.Server.Port),
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return nil
}

func (rs *HttpService) Start() {
	conf := rs.cs.GetConfig()

	go func() {
		if err := rs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
	}()

	log.Println("server started on port " + fmt.Sprintf("%d", conf.Server.Port))
}

func (rs *HttpService) Shutdown(ctx context.Context) {
	if err := rs.server.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}

func (rs *HttpService) GetRoute() *mux.Router {
	return rs.route
}
