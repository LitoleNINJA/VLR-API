package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type MatchStatus string

const (
	Upcoming  MatchStatus = "upcoming"
	Live      MatchStatus = "live"
	Completed MatchStatus = "completed"
)

type Match struct {
	URL       string
	Team1     string
	Team2     string
	Score     []int
	Rounds    []int
	StartTime string
	Tag       string
	Status    MatchStatus
	Region    string
}

func GetMatchesFromVLR() []Match {
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

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	return matches
}

func GetResultsFromVLR() []Match {
	c := colly.NewCollector()

	var matches []Match
	c.OnHTML("div.wf-card", func(e *colly.HTMLElement) {
		e.ForEach("a.wf-module-item", func(_ int, el *colly.HTMLElement) {
			match, err := getMatchResult(el)
			if err != nil {
				fmt.Println(err)
				return
			}
			matches = append(matches, match)
		})
	})

	c.Visit("https://www.vlr.gg/matches/results")

	c.Wait()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	return matches
}
