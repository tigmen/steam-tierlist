package main

import(
	"github.com/BurntSushi/toml"
	"github.com/tigmen/steam-tierlist/internal/loader"
	"flag"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, 
	"configPath", "configs/loader.toml",
	"path to config file")
}

func main() {
	l := loader.NewLoader()
	
	_, err := toml.DecodeFile(configPath, l.Config)
	if err != nil {

	}
}
