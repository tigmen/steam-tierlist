package loader

type Config struct {
	Key      string `toml:"key"`
	Userid   string `toml:"userid"`
	LogLevel string `toml:"log_level"`
	OutFilePath  string `toml:outfilepath"`
}

func NewConfig() *Config {
	return &Config{}
}
