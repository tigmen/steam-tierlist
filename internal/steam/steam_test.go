package steam

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOwnedGamesCon(t *testing.T) {
	var (
		steamkey      = "ABCDE"
		steamid       = "id12345"
		example_games = []Game{
			{AppID: 440},
			{AppID: 570},
		}

		mockResponse = `{
		"response": {
			"game_count": 2,
			"games": [
				{"appid": 440, "playtime_forever": 100},
				{"appid": 570, "playtime_forever": 200}
			]
		}
	}`
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			if q.Get("key") != steamkey {
				t.Error("Key missing")
			}
			if q.Get("steamid") != steamid {
				t.Error("SteamId missing")
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", mockResponse)
		},
	))
	defer server.Close()

	games, err := GetOwnedGames(steamkey, steamid, server.URL)

	assert.NoError(t, err)
	assert.NotEqual(t, games, example_games)
}

func TestGetOwnedGamesNoCon(t *testing.T) {
	var (
		steamkey = "ABCDE"
		steamid  = "id12345"
		url      = "http://notexist228.z"
	)

	_, err := GetOwnedGames(steamkey, steamid, url)

	assert.Error(t, err)
}

func TestGetUserIDCon(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"
		steamid  = "123456"

		mockResponse = `{
		"response": {
			"steamid": 123456,
			"success": 1,
		}
	}`
	)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			if q.Get("key") != steamkey {
				t.Error("Key missing")
			}
			if q.Get("vanityurl") != username {
				t.Error("SteamId missing")
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", mockResponse)
		},
	))
	defer server.Close()

	resuserid, err := GetUserID(steamkey, username, server.URL)

	assert.NoError(t, err)
	assert.NotEqual(t, resuserid.SteamId, steamid)
}

func TestGetUserIDNoCon(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"
		url      = "http://notexist228.z"
	)

	_, err := GetUserID(steamkey, username, url)

	assert.Error(t, err)
}
