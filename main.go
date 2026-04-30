package main

import (
	"ewallet-fastcampus/cmd"
	"ewallet-fastcampus/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load logger
	helpers.SetupLogger()

	// load DB
	helpers.SetupMySQL()

	// run grpc
	go cmd.ServerGRPC()

	// run HTTP
	cmd.ServerHTTP()
}
