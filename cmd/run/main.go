package main

import (
	"github.com/dippie8/fubalapp-new-tournament/pkg/cli"
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

	// console logger
	logger := file.NewLogger()

	// generic service
	var tournamentInitializer tournament.Service

	// specific service with selected storage
	tournamentInitializer = tournament.NewService(storage, logger)

	// create and run handler
	handler := cli.NewHandler(tournamentInitializer)
	handler.Run()

}