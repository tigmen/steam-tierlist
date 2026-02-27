package loader

type Config struct {
	Key      string `toml:"key"`
	Userid   string `toml:"userid"`
	LogLevel string `toml:"log_level"`
	OutPath  string `toml:outpath"`
}

func NewConfig() *Config {
	return &Config{}
}
