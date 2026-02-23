package loader

type Config struct {
	key string `toml:"key"`
	userid int `toml:"userid"`
}

func NewConfig() * Config {
	return &Config{
	}
}
