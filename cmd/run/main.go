package main

import (
	"github.com/dippie8/fubalapp-new-tournament/pkg/presenter/file"
	"github.com/dippie8/fubalapp-new-tournament/pkg/storage/mongodb"
	"github.com/dippie8/fubalapp-new-tournament/pkg/tournament"
)

func main() {

	// Generic storage
	var storage tournament.Repository

	// MongoDB storage
	storage, err := mongodb.NewDB()
	defer storage.Disconnect()

	if err != nil {
		panic(err)
	}

	// file logger
	logger := file.NewLogger()

	// abstract service
	var tournamentInitializer tournament.Service

	// specific service with selected storage and logger
	tournamentInitializer = tournament.NewService(storage, logger)

	err = tournamentInitializer.StartNewTournament()
	if err != nil {
		panic(err)
	}

}