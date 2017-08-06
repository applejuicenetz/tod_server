package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	SetBasicLogger()
	log.Info("Starting TOD server")
	config_flag := flag.String("config", "", "Path to config file")
	flag.Parse()

	config := DefaultConfig()
	if *config_flag != "" {
		var err error
		config, err = LoadConfig(*config_flag)
		if err != nil {
			log.Error("Error loading config file: %v. Using default config", err)
			config = DefaultConfig()
		}
	}

	ResetLogging(config)

	if !config.TCPEnabled && !config.UDPEnabled {
		log.Fatalf("At least one of tcp or udp must be enabled")
	}

	tcp_server := TCPServer{}
	go tcp_server.Run(config)

	udp_server := UDPServer{}
	go udp_server.Run(config)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Info("Exiting TOD server")
}
