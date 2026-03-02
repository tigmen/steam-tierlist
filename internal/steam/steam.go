package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Game struct {
	AppID int64 `json:"appid"`
}

type responseOwnedGames struct {
	Game_count int64  `json:"game_count"`
	Games      []Game `json:"games"`
}

func GetOwnedGames(key, steamid string, url string) (*[]Game, error) {
	res, err := http.Get(fmt.Sprintf(
		"%s/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json",
		url, key, steamid))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	httpresponse := make(map[string]responseOwnedGames)

	json.NewDecoder(res.Body).Decode(&httpresponse)

	out := httpresponse["response"].Games
	return &out, nil
}

type responseUserId struct {
	SteamId string `json:"steamid"`
	Success int64  `json:"success"`
}

func GetUserID(key, nickname string, url string) (*responseUserId, error) {
	res, err := http.Get(fmt.Sprintf(
		"%s/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s",
		url, key, nickname))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	httpresponse := make(map[string]responseUserId)

	json.NewDecoder(res.Body).Decode(&httpresponse)

	out := httpresponse["response"]
	return &out, nil
}
