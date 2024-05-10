package scrapper

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func getMatch(e *colly.HTMLElement) (Match, error) {
	match := Match{}

	match.StartTime = strings.TrimSpace(e.ChildText("div.h-match-eta"))
	match.Tag = strings.TrimSpace(e.ChildText("div.h-match-preview-event"))
	match.Status = Upcoming
	match.Score = []int{0, 0}

	e.ForEach("div.h-match-team-name", func(_ int, el *colly.HTMLElement) {
		if match.Team1 == "" {
			match.Team1 = strings.TrimSpace(el.Text)
		} else {
			match.Team2 = strings.TrimSpace(el.Text)
		}
	})

	log.Println(match)

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

	log.Println(matches)
	return nil
}
