package initializing

import (
	"sort"
)

type Medal int

const (
	Gold Medal = iota
	Silver
	Bronze

	minGamesEligibility = 5
)

type Repository interface {
	// GetStandings returns the list of player with number of wins and Played games
	GetStandings() ([]Standing, error)
	// AddPrize assign a medal to a player
	AddPrize(string, Medal) error
	// ResetStandings reset wins and Played of everyone
	ResetStandings() error
	// Disconnect from repository
	Disconnect()
}

type Service interface {
	RewardPlayers() (Prizes, error)
	ResetLeague() error
	//Initialize() (Prizes, error)
}

type service struct {
	r Repository
}

type Prizes struct {
	First []string
	Second []string
	Third []string
}

// NewService create a initializing service with dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// Initialize rewards players and reset standings
func (s service) RewardPlayers() (prizes Prizes ,err error) {

	standings, err := s.r.GetStandings()

	//standings = sortStandings(standings)
	topRates := topThreeRates(standings)
	prizes.First = []string{}
	prizes.Second = []string{}
	prizes.Third = []string{}

	for _, player := range standings {

		if player.Played >= minGamesEligibility {
			winRate := float32(player.Win) / float32(player.Played)

			switch winRate {
			case topRates[0]:
				prizes.First = append(prizes.First, player.Id)
				err = s.r.AddPrize(player.Id, Gold)
			case topRates[1]:
				prizes.Second = append(prizes.Second, player.Id)
				err = s.r.AddPrize(player.Id, Silver)
			case topRates[2]:
				prizes.Third = append(prizes.Third, player.Id)
				err = s.r.AddPrize(player.Id, Bronze)
			}
		}
	}

	return prizes, err
}

func (s service) ResetLeague() error {
	return s.r.ResetStandings()
}

// topThreeRates return top 3 values of winning rate
func topThreeRates(standings []Standing) (topRates [3]float32) {
	keys := make(map[float32]bool)
	var rates []float32
	for _, player := range standings {

		if player.Played >= minGamesEligibility {
			rate := float32(player.Win) / float32(player.Played)

			if _, value := keys[rate]; !value {
				keys[rate] = true
				rates = append(rates, rate)
			}
		}
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i] > rates[j]
	})

	for i := 0; i < 3 && i < len(rates); i++ {
		topRates[i] = rates[i]
	}

	return topRates
}


