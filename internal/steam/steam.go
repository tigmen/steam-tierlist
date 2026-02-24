package steam

import (
	"fmt"
	"net/http"
)

func GetOwnedGames(key, steamid string) (http.Response, error) {
	res, err := http.Get(fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json"))
	if err != nil {
		return nil, err
	}

	return res.Body, err
}
