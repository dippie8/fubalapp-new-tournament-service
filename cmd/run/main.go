package main

import (
	"github.com/dippie8/fubalapp-new-tournament/pkg/presenter/file"
	"github.com/dippie8/fubalapp-new-tournament/pkg/storage/mongodb"
	"github.com/dippie8/fubalapp-new-tournament/pkg/tournament"
)

func main() {

	// file logger
	logger := file.NewLogger()

	// MongoDB storage
	storage, err := mongodb.NewDB()
	defer storage.Disconnect()

	if err != nil {
		logger.Log(err.Error())
		panic(err)
	}

	// new tournament service
	var tournamentInitializer tournament.Service
	tournamentInitializer = tournament.NewService(storage, logger)

	// start new tournament
	err = tournamentInitializer.StartNewTournament()
	if err != nil {
		logger.Log(err.Error())
		panic(err)
	}
}