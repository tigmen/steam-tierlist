package steam

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetOwnedGames_Con(t *testing.T) {
	var (
		steamkey      = "ABCDE"
		steamid       = "id12345"
		example_games = &[]Game{
			{AppID: 440},
			{AppID: 570},
		}

		mockResponse = `{
		"response": {
			"success": 1,
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
	assert.Equal(t, games, example_games)
}

func Test_GetOwnedGames_IncorrectJSON(t *testing.T) {
	var (
		steamkey     = "ABCDE"
		steamid      = "id12345"
		mockResponse = `{
		"response": {
			"success": 1,
			"game_count": 2,
			"games": [
				{"appid": 440, "playtime_forever": 100},
				{"appid": 570, "playtime_forever": 200}
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

	_, err := GetOwnedGames(steamkey, steamid, server.URL)

	assert.Error(t, err)
}

func Test_GetOwnedGames_Unsuccesful(t *testing.T) {
	var (
		steamkey     = "ABCDE"
		steamid      = "id12345"
		mockResponse = `{
		"response": {
			"success": 0,
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

	_, err := GetOwnedGames(steamkey, steamid, server.URL)

	assert.Error(t, err)
}
func Test_GetOwnedGames_NoCon(t *testing.T) {
	var (
		steamkey = "ABCDE"
		steamid  = "id12345"
		url      = "http://notexist228.z"
	)

	_, err := GetOwnedGames(steamkey, steamid, url)

	assert.Error(t, err)
}

func Test_GetUserID_Con(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"
		steamid  = "123456"

		mockResponse = `{
			"response": {
				"success": 1,
				"steamid": "123456"
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
	assert.Equal(t, resuserid.SteamId, steamid)
}

func Test_GetUserID_IncorrectJSON(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"

		mockResponse = `{
			"response": {
				"success": 1,
				"steamid": 123456,
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

	_, err := GetUserID(steamkey, username, server.URL)

	assert.Error(t, err)
}

func Test_GetUserID_Unsuccesful(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"

		mockResponse = `{
			"response": {
				"success": 0,
				"steamid": "123456"	
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

	_, err := GetUserID(steamkey, username, server.URL)

	assert.Error(t, err)
}

func Test_GetUserID_NoCon(t *testing.T) {
	var (
		steamkey = "ABCDE"
		username = "testuser"
		url      = "http://notexist228.z"
	)

	_, err := GetUserID(steamkey, username, url)

	assert.Error(t, err)
}
