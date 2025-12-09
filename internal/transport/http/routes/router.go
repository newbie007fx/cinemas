package routes

import (
	"net/http"

	"github.com/newbie007fx/cinemas/internal/dependencies"
	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/transport/http/handler/authentication"
	"github.com/newbie007fx/cinemas/internal/transport/http/handler/showtimes"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/response"
	"github.com/newbie007fx/cinemas/internal/transport/http/middleware"
	"github.com/newbie007fx/cinemas/platform/httpserver"

	"github.com/gorilla/mux"
)

type apiRouter struct {
	dep              *dependencies.Dependency
	baseRoute        *mux.Router
	authController   *authentication.AuthHandlers
	showtimeHandlers *showtimes.Handler
}

func Init(httpService *httpserver.HttpService, dep *dependencies.Dependency) {
	router := &apiRouter{
		baseRoute:        httpService.GetRoute(),
		dep:              dep,
		authController:   authentication.New(dep.UserUsecase, dep.AuthToken),
		showtimeHandlers: showtimes.New(dep.ShowtimeUsecase),
	}

	router.baseRoute.HandleFunc("/ping", router.ping).Methods(http.MethodGet)
	router.baseRoute.NotFoundHandler = http.HandlerFunc(router.notFound)

	router.initApiv1()

}

func (ar apiRouter) initApiv1() {
	loginRoute := ar.baseRoute.Path("/api/v1/auth/login").Subrouter()
	loginRoute.HandleFunc("", ar.authController.Login).Methods(http.MethodPost)

	apiv1Route := ar.baseRoute.PathPrefix("/api/v1").Subrouter()

	apiv1Route.Use(middleware.Auth(ar.dep.AuthToken))

	showtimeRoute := apiv1Route.PathPrefix("/showtimes").Subrouter()
	showtimeRoute.HandleFunc("", ar.showtimeHandlers.List).Methods(http.MethodGet)
	showtimeRoute.HandleFunc("", ar.showtimeHandlers.Create).Methods(http.MethodPost)
	showtimeRoute.HandleFunc("/{id}", ar.showtimeHandlers.Get).Methods(http.MethodGet)
	showtimeRoute.HandleFunc("/{id}", ar.showtimeHandlers.Update).Methods(http.MethodPut)
	showtimeRoute.HandleFunc("/{id}", ar.showtimeHandlers.Delete).Methods(http.MethodDelete)
}

func (apiRouter) notFound(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[map[string]string]{
		Error: errors.ErrorPathNotFound.New("path not found"),
	}

	resp.Send(w)
}

func (apiRouter) ping(w http.ResponseWriter, r *http.Request) {
	resp := response.Response[map[string]string]{
		IsSuccess: true,
		Data:      map[string]string{"ping": "Pong!!!"},
	}

	resp.Send(w)
}
