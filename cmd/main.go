package main

import (
	"github.com/SifatU360/GoTix-.git/internal/config"
	"github.com/SifatU360/GoTix-.git/internal/server"
)

func main() {

	cfg := config.LoadEnv()

	db := config.ConnectDatabase(cfg)

	server.Start(db, cfg)

}
