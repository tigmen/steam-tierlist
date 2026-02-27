package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/tigmen/steam-tierlist/internal/loader"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/loader.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := loader.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	l, err := loader.NewLoader(config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := l.Load(); err != nil {
		logrus.Fatal(err)
	}
}
