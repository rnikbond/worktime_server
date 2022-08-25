package handler

import (
	"worktime_server/internal/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Post("/login", h.Login)

	router.Route("/ticker", func(r chi.Router) {
		r.Post("/start", h.TickerStart)
		r.Post("/stop", h.TickerStop)
	})

	return router
}
