package scrapper

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lithammer/shortuuid"
)

func getMatch(e *colly.HTMLElement) (Match, error) {
	match := Match{}

	match.ID = shortuuid.New()
	match.StartTime = strings.TrimSpace(e.ChildText("div.h-match-eta"))
	match.Tag = strings.TrimSpace(e.ChildText("div.h-match-preview-event"))
	match.Status = Upcoming
	match.Score = []int{0, 0}
	match.Rounds = []int{0, 0}

	e.ForEach("div.h-match-team", func(_ int, el *colly.HTMLElement) {
		if match.Team1 == "" {
			match.Team1 = strings.TrimSpace(el.ChildText("div.h-match-team-name"))
			match.Score[0], _ = strconv.Atoi(el.ChildText("div.h-match-team-score"))
			match.Rounds[0], _ = strconv.Atoi(el.ChildText("div.h-match-team-rounds > span"))
		} else {
			match.Team2 = strings.TrimSpace(el.ChildText("div.h-match-team-name"))
			match.Score[1], _ = strconv.Atoi(el.ChildText("div.h-match-team-score"))
			match.Rounds[1], _ = strconv.Atoi(el.ChildText("div.h-match-team-rounds > span"))
		}

	})

	return match, nil
}

func saveMatchData(matches []Match) error {
	matchData, err := json.Marshal(matches)
	if err != nil {
		return err
	}

	matchFile, err := os.OpenFile("matches.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer matchFile.Close()

	_, err = matchFile.Write(matchData)
	if err != nil {
		return err
	}

	return nil
}

func MarshalMatches(matches []Match) string {
	matchData, _ := json.Marshal(matches)
	return string(matchData)
}

func UnmarshalMatches(data string) []Match {
	var matches []Match
	json.Unmarshal([]byte(data), &matches)
	return matches
}
