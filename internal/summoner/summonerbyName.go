package summoner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cantevenread/kup"
)

// GetSummonersByName needs a name and returns a summoner type. can error.
func GetSummonersByName(name string, c kup.KupClient) (*Summoner, error) {
	url := "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + name + "?api_key=" + c.RGAPI

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var summoner Summoner
	if err := json.Unmarshal(body, &summoner); err != nil {
		return nil, err
	}

	return &summoner, nil
}
