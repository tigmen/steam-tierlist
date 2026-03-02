package loader

type Config struct {
	Key      string `toml:"key"`
	Userid   string `toml:"userid"`
	LogLevel string `toml:"log_level"`
	OutPath  string `toml:"outpath"`
	UserName string `toml:"username"`
}

func NewConfig() *Config {
	return &Config{}
}
