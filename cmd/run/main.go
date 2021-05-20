package main

import (
	"github.com/dippie8/fubalapp-new-tournament/pkg/cli"
	"github.com/dippie8/fubalapp-new-tournament/pkg/initializing"
	"github.com/dippie8/fubalapp-new-tournament/pkg/storage/mongodb"
)

func main() {

	// Generic storage
	var storage initializing.Repository

	// MongoDB storage
	storage, err := mongodb.NewDB()
	defer storage.Disconnect()

	if err != nil {
		panic(err)
	}

	// generic service
	var leagueInitializer initializing.Service

	// specific service with selected storage
	leagueInitializer = initializing.NewService(storage)

	// create and run handler
	handler := cli.NewHandler(leagueInitializer)
	handler.Run()

}