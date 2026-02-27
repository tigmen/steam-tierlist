package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Game struct {
	AppID int64 `json:"appid"`
}

type response struct {
	Game_count int64  `json:"game_count"`
	Games      []Game `json:"games"`
}
type httpResponse struct {
	response `json:"response"`
}

func GetOwnedGames(key, steamid string) (*[]string, error) {
	res, err := http.Get(fmt.Sprintf(
		"http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json",
		key, steamid))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	httpresponse := &httpResponse{}

	json.NewDecoder(res.Body).Decode(httpresponse)

	imageurl := make([]string, 0)

	for _, game := range httpresponse.response.Games {
		imageurl = append(imageurl,
			fmt.Sprintf("https://cdn.akamai.steamstatic.com/steam/apps/%d/library_600x900.jpg", game.AppID))
	}

	return &imageurl, nil
}
