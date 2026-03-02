package loader

import (
	"bufio"
	"fmt"
	"os"

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
	userid := l.config.Userid 
	if userid == "" {
		res_userid, err := steam.GetUserID(l.config.Key, l.config.UserName)
		if err != nil {
			return err
		}
		userid = res_userid.SteamId
	}

	res, err := steam.GetOwnedGames(l.config.Key, userid)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/data.js", l.config.OutPath) , os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.Write([]byte("const steamGames = [\n"))

	for _, game := range *res {
		_, err := fmt.Fprintf(writer, "\t{ id: \"%d\", name: \"%s\" },\n", game.AppID, "unit")
		if err != nil {
			return nil
		}
	}

	writer.Write([]byte("]"))
	err = writer.Flush()
	if err != nil {
		return err
	}

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
