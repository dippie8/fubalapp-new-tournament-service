package cli

import (
	"fmt"
	"github.com/dippie8/fubalapp-new-tournament/pkg/initializing"
	"time"
)

type Handler struct {
	initializingService initializing.Service
}

func NewHandler(service initializing.Service) *Handler {
	return &Handler{
		initializingService: service,
	}
}

func (h *Handler) Run() {
	rewarded, err := h.initializingService.RewardPlayers()
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().String() + ":")
	fmt.Println("gold medals to:", rewarded.First)
	fmt.Println("silver medals to:", rewarded.Second)
	fmt.Println("bronze medals to:", rewarded.Third)
	fmt.Print("------------------------\n\n")

	err = h.initializingService.ResetLeague()

	if err != nil {
		panic(err)
	}
}

