package cli

import (
	"github.com/dippie8/fubalapp-new-tournament/pkg/tournament"
)

type Handler struct {
	initializingService tournament.Service
}

func NewHandler(service tournament.Service) *Handler {
	return &Handler{
		initializingService: service,
	}
}

func (h *Handler) Run() {
	err := h.initializingService.StartNewTournament()
	if err != nil {
		panic(err)
	}
}

