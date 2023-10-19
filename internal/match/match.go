package match

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cantevenread/kup"
	"github.com/cantevenread/kup/internal/summoner"
)

type Match struct {
	MatchID string `json:"0"`
	Info    MatchInfo
}

type MatchInfo struct {
	GameDuration int      `json:"gameDuration"`
	MapID        int      `json:"mapId"`
	Participants []string `json:"metadata"`
	// Add more fields for match info as needed
}

func FindLatestMatchPlayed(s summoner.Summoner, c kup.KupClient) (*Match, error) {
	url := strings.ToLower(string(c.Server)) + "/lol/match/v5/matches/by-puuid/" + s.PullPUUID() + "/ids" + "?start=0&count=1&api_key=" + c.RGAPI
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	var matchIDs []string
	if err := json.NewDecoder(response.Body).Decode(&matchIDs); err != nil {
		return nil, err
	}

	if len(matchIDs) == 0 {
		return nil, fmt.Errorf("No match IDs found in the JSON response")
	}

	// Create a Match struct with the first match ID in the array
	match := &Match{
		MatchID: matchIDs[0],
	}

	return match, nil
}

func FetchMatchInfo(match *Match, c kup.KupClient) (MatchInfo, error) {
	url := strings.ToLower(string(c.Server)) + "/lol/match/v5/matches/" + match.MatchID + "?api_key=" + c.RGAPI
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return MatchInfo{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return MatchInfo{}, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	var info MatchInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return MatchInfo{}, err
	}

	return info, nil
}
