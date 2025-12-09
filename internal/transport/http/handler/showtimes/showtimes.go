package showtimes

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/newbie007fx/cinemas/internal/errors"
	showtimeModels "github.com/newbie007fx/cinemas/internal/module/showtimes/models"
	showtimePorts "github.com/newbie007fx/cinemas/internal/module/showtimes/ports"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/request"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/response"
)

type Handler struct {
	usecase showtimePorts.Usecase
}

func New(usecase showtimePorts.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload showtimePayload
	if err := request.ReadRequestBody(r, &payload); err != nil {
		response.WriterResponseError(w, err)
		return
	}

	showtime, err := h.usecase.Create(r.Context(), payload.toCreateInput())
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	resp := response.Response[*showtimeModels.Showtime]{
		IsSuccess: true,
		Data:      showtime,
	}

	resp.Send(w)
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	showtimes, err := h.usecase.List(r.Context())
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	resp := response.Response[[]showtimeModels.Showtime]{
		IsSuccess: true,
		Data:      showtimes,
	}

	resp.Send(w)
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r)
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	showtime, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	resp := response.Response[*showtimeModels.Showtime]{
		IsSuccess: true,
		Data:      showtime,
	}

	resp.Send(w)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r)
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	var payload showtimePayload
	if err := request.ReadRequestBody(r, &payload); err != nil {
		response.WriterResponseError(w, err)
		return
	}

	showtime, err := h.usecase.Update(r.Context(), id, payload.toUpdateInput())
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	resp := response.Response[*showtimeModels.Showtime]{
		IsSuccess: true,
		Data:      showtime,
	}

	resp.Send(w)
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r)
	if err != nil {
		response.WriterResponseError(w, err)
		return
	}

	if err := h.usecase.Delete(r.Context(), id); err != nil {
		response.WriterResponseError(w, err)
		return
	}

	resp := response.Response[map[string]string]{
		IsSuccess: true,
		Data:      map[string]string{"message": "showtime deleted"},
	}

	resp.Send(w)
}

func parseIDFromPath(r *http.Request) (uint, *errors.BaseError) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		return 0, errors.ErrorInvalidPathValue.New("missing showtime id")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errors.ErrorInvalidPathValue.New("invalid showtime id")
	}

	return uint(id), nil
}
