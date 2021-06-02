package tournament

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockStorage struct {
	standings []Standing
	players []Player
}

type mockLogger struct {}

func (m *mockStorage) Disconnect() {
	fmt.Println("disconnected")
}

func (m *mockStorage) GetStandings() ([]Standing, error) {
	return m.standings, nil
}

func (m *mockStorage) AddPrize(name string, medal Medal) error {
	for i, _ := range m.players {
		if m.players[i].name == name {
			switch medal {
			case Gold:
				m.players[i].gold ++
			case Silver:
				m.players[i].silver ++
			case Bronze:
				m.players[i].silver ++
			}
		}
	}
	return nil
}

// ResetStandings reset wins and Played of everyone
func (m *mockStorage)  ResetStandings() error {
	for i, _ := range m.standings {
		m.standings[i].Played = 0
		m.standings[i].Win = 0
	}
	return nil
}

func (m *mockLogger) Log(message string) {
	datetime := time.Now().String()
	fmt.Println(datetime + ": " + message)
}

func Test_service_rewards_standard(t *testing.T) {

	standings := []Standing{
		{
			Id:     "dippi",
			Win:    6,
			Played: 8,
		},
		{
			Id:     "moro",
			Win:    4,
			Played: 8,
		},
		{
			Id:     "fra",
			Win:    9,
			Played: 12,
		},
		{
			Id:     "angelo",
			Win:    3,
			Played: 8,
		},
		{
			Id:     "bardo",
			Win:    6,
			Played: 16,
		},
		{
			Id:     "teo",
			Win:    2,
			Played: 8,
		},
	}

	logger := new(mockLogger)
	storage := new(mockStorage)
	storage.standings = standings

	service := NewService(storage, logger)

	prizes, err := service.RewardPlayers()

	if err != nil {
		t.Error(err)
	}

	goldList := []string{"dippi","fra"}
	silverList := []string{"moro"}
	bronzeList := []string{"angelo","bardo"}

	assert.Equal(t, goldList, prizes.First)
	assert.Equal(t, silverList, prizes.Second)
	assert.Equal(t, bronzeList, prizes.Third)
}

func Test_service_rewards_hybrid(t *testing.T) {

	standings := []Standing{
		{
			Id:     "dippi",
			Win:    3,
			Played: 4,
		},
		{
			Id:     "moro",
			Win:    8,
			Played: 8,
		},
		{
			Id:     "fra",
			Win:    0,
			Played: 8,
		},
		{
			Id:     "angelo",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "bardo",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "teo",
			Win:    0,
			Played: 0,
		},
	}

	logger := new(mockLogger)
	storage := new(mockStorage)
	storage.standings = standings

	service := NewService(storage, logger)

	prizes, err := service.RewardPlayers()

	if err != nil {
		t.Error(err)
	}

	goldList := []string{"moro"}
	silverList := []string{"fra"}
	bronzeList := []string{}

	assert.Equal(t, goldList, prizes.First)
	assert.Equal(t, silverList, prizes.Second)
	assert.Equal(t, bronzeList, prizes.Third)
}

func Test_service_rewards_noGames(t *testing.T) {

	standings := []Standing{
		{
			Id:     "dippi",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "moro",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "fra",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "angelo",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "bardo",
			Win:    0,
			Played: 0,
		},
		{
			Id:     "teo",
			Win:    0,
			Played: 0,
		},
	}

	logger := new(mockLogger)
	storage := new(mockStorage)
	storage.standings = standings

	service := NewService(storage, logger)

	prizes, err := service.RewardPlayers()

	if err != nil {
		t.Error(err)
	}

	goldList := []string{}
	silverList := []string{}
	bronzeList := []string{}

	assert.Equal(t, goldList, prizes.First)
	assert.Equal(t, silverList, prizes.Second)
	assert.Equal(t, bronzeList, prizes.Third)
}

func Test_service_ResetLeague(t *testing.T) {
	standings := []Standing{
		{
			Id:     "dippi",
			Win:    6,
			Played: 8,
		},
		{
			Id:     "moro",
			Win:    4,
			Played: 8,
		},
		{
			Id:     "fra",
			Win:    3,
			Played: 4,
		},
		{
			Id:     "angelo",
			Win:    3,
			Played: 8,
		},
		{
			Id:     "bardo",
			Win:    6,
			Played: 16,
		},
		{
			Id:     "teo",
			Win:    2,
			Played: 8,
		},
	}
	storage := &mockStorage{standings: standings}
	logger := new(mockLogger)

	service := NewService(storage, logger)

	err := service.ResetLeague()

	if err != nil {
		t.Error(err)
	}

	newStandings, err := storage.GetStandings()
	if err != nil {
		t.Error(err)
	}

	for _, st := range newStandings{
		assert.Equal(t, 0, st.Win)
		assert.Equal(t, 0, st.Played)
	}
}

func Test_service_StartNewTournament(t *testing.T) {
	standings := []Standing{
		{
			Id:     "dippi",
			Win:    6,
			Played: 8,
		},
		{
			Id:     "moro",
			Win:    4,
			Played: 8,
		},
		{
			Id:     "fra",
			Win:    3,
			Played: 4,
		},
		{
			Id:     "angelo",
			Win:    3,
			Played: 8,
		},
		{
			Id:     "bardo",
			Win:    6,
			Played: 16,
		},
		{
			Id:     "teo",
			Win:    2,
			Played: 8,
		},
	}
	storage := &mockStorage{standings: standings}
	logger := new(mockLogger)

	service := NewService(storage, logger)

	err:= service.StartNewTournament()

	if err != nil {
		t.Error(err)
	}

	newStandings, err := storage.GetStandings()
	if err != nil {
		t.Error(err)
	}

	for _, st := range newStandings {
		assert.Equal(t, 0, st.Win)
		assert.Equal(t, 0, st.Played)
	}

}