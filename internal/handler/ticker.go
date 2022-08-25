package handler

import (
	"fmt"
	"net/http"

	"worktime_server/internal/service"
)

func (h *Handler) TickerStart(w http.ResponseWriter, r *http.Request) {

	fmt.Println("endpoint :: TickerStart")

	if err := h.services.RunTasks("admin"); err != nil {

		fmt.Printf("failed start ticker: %v\n", err)
		w.WriteHeader(service.ToHTTP(err))
		return
	}
}

func (h *Handler) TickerStop(w http.ResponseWriter, r *http.Request) {

	fmt.Println("endpoint :: TickerStop")

	if err := h.services.StopTasks("admin"); err != nil {

		fmt.Printf("failed stop ticker: %v\n", err)
		w.WriteHeader(service.ToHTTP(err))
		return
	}
}
