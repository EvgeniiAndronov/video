package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"video/iternal/app/videoserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/videoserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := videoserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := videoserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
