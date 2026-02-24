package loader

import (
	"github.com/sirupsen/logrus"
	"github.com/tigmen/steam-tierlist/internal/steam"
)

type Loader struct {
	logger *logrus.Logger
	config *Config
}

func NewLoader(config *Config) (*Loader, error) {
	l := &Loader{
		config: config,
		logger: logrus.New(),
	}

	err := l.configureLogger()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l Loader) Load() error {
	gamesReader, err := steam.GetOwnedGames(l.config.Key, l.config.Userid)
	if err != nil {
		return err
	}

	gamesReader.Read(buffer)

	l.logger.Info()

	return nil
}

func (l *Loader) configureLogger() error {
	level, err := logrus.ParseLevel(l.config.LogLevel)
	if err != nil {
		return err
	}
	l.logger.SetFormatter(&logrus.JSONFormatter{})
	l.logger.SetLevel(level)

	return nil
}
