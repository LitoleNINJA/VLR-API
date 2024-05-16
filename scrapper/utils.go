package scrapper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var regions = map[string]string{
	"Americas": "NA",
	"EMEA":     "EU",
	"Pacific":  "APAC",
	"China":    "CN",
	"Brazil":   "BR",
	"Spain":    "ES",
	"Japan":    "JP",
	"Korea":    "KR",
	"LATAM":    "LATAM",
}

func getMatch(e *colly.HTMLElement) (Match, error) {
	match := Match{}

	if e.Request == nil {
		return match, fmt.Errorf("request is nil")
	}

	match.URL = e.Request.AbsoluteURL(e.Attr("href"))
	match.StartTime = strings.TrimSpace(e.ChildText("div.h-match-eta"))
	match.Tag = strings.TrimSpace(e.ChildText("div.h-match-preview-event"))
	if match.StartTime == "LIVE" {
		match.Status = Live
	} else {
		match.Status = Upcoming
	}
	match.Score = []int{0, 0}
	match.Rounds = []int{0, 0}
	match.Region = findRegion(match.Tag)

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

func getMatchResult(e *colly.HTMLElement) (Match, error) {
	match := Match{}

	match.URL = e.Request.AbsoluteURL(e.Attr("href"))
	match.Tag = strings.TrimSpace(e.ChildText("div.match-item-event.text-of"))
	match.Status = Completed
	match.Score = []int{0, 0}
	match.StartTime = strings.TrimSpace(e.ChildText("div.ml-eta.mod-completed"))
	match.Region = findRegion(match.Tag)

	e.ForEach("div.match-item-vs-team", func(_ int, el *colly.HTMLElement) {
		if match.Team1 == "" {
			match.Team1 = strings.TrimSpace(el.ChildText("div.match-item-vs-team-name > div.text-of"))
			match.Score[0], _ = strconv.Atoi(el.ChildText("div.match-item-vs-team-score"))
		} else {
			match.Team2 = strings.TrimSpace(el.ChildText("div.match-item-vs-team-name > div.text-of"))
			match.Score[1], _ = strconv.Atoi(el.ChildText("div.match-item-vs-team-score"))
		}
	})

	return match, nil
}

func findRegion(tag string) string {
	for region, code := range regions {
		if strings.Contains(tag, region) {
			return code
		}
	}
	return "Unknown"
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
