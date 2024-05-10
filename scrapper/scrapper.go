package scrapper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type MatchStatus string

const (
	Upcoming  MatchStatus = "upcoming"
	Live      MatchStatus = "live"
	Completed MatchStatus = "completed"
)

type Match struct {
	Team1     string
	Team2     string
	Score     []int
	StartTime string
	Tag       string
	Status    MatchStatus
}

func GetMatches() []Match {
	c := colly.NewCollector()

	var matches []Match
	c.OnHTML("div.js-home-matches-upcoming", func(e *colly.HTMLElement) {
		e.ForEach("a.wf-module-item", func(_ int, el *colly.HTMLElement) {
			match, err := getMatch(el)
			if err != nil {
				fmt.Println(err)
				return
			}
			matches = append(matches, match)
		})
	})

	c.Visit("https://www.vlr.gg")

	c.Wait()

	saveMatchData(matches)

	logfile, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)

	return matches
}
