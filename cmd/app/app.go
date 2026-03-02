package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/tigmen/steam-tierlist/internal/loader"
)

var (
	configPath string
	username string
	outPath    string
	userId     string
	log_level  string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/loader.toml", "path to config file")
	flag.StringVar(&username, "username", "", "steam username from url")
	flag.StringVar(&outPath, "o", "bin", "Path to output files. Must consists index.html")
	flag.StringVar(&userId, "userid", "", "User ID to load")
	flag.StringVar(&log_level, "log_level", "debug", "Logging level")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	flag.Parse()

	config := loader.NewConfig()

	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	config.Key = os.Getenv("STEAMAPIKEY")
	config.Userid = userId
	config.OutPath = outPath
	config.LogLevel = log_level
	config.UserName = username

	l, err := loader.NewLoader(config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := l.Load(); err != nil {
		logrus.Fatal(err)
	}
}
