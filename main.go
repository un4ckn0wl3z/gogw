package main

import ( 
	"flag"

	"gogw/logger"
	"gogw/config"
	"gogw/server"
	"gogw/client"
)

var cfgFile = flag.String("c", "cfg.json", "")

func main(){
	logger.LEVEL = logger.INFO

	logger.Info("gogw start")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgFile)
	if err != nil {
		logger.Error(err)
		return
	}

	if cfg.Role == "server" {
		server := server.NewServer(cfg.Server.ServerAddr)
		server.Start()
	}

	if cfg.Role == "client" {
		client := client.NewClient(cfg.Client.ServerAddr, cfg.Client.LocalAddr, cfg.Client.RemotePort)
		client.Start()
	}
}