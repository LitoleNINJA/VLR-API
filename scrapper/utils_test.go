package scrapper

import (
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
)

func TestGetMatch(t *testing.T) {
	e, err := getCollyHTMLElement()
	assert.NoError(t, err, "getCollyHTMLElement failed")

	match, err := getMatch(&e)
	assert.NoError(t, err, "getMatch failed")

	assert.NotNil(t, match, "getMatch returned nil")
	assert.NotNil(t, match.URL, "URL is nil")
	assert.NotNil(t, match.Team1, "Team1 is nil")
	assert.NotNil(t, match.Team2, "Team2 is nil")
	assert.NotNil(t, match.Score, "Score is nil")
	assert.NotNil(t, match.Rounds, "Rounds is nil")
	assert.NotNil(t, match.StartTime, "StartTime is nil")
	assert.NotNil(t, match.Tag, "Tag is nil")
	assert.NotNil(t, match.Status, "Status is nil")
	assert.NotNil(t, match.Region, "Region is nil")

	assert.NotEqual(t, match.Status, Completed, "Status is Completed, expected Upcoming or Live")
}

func TestGetMatchResult(t *testing.T) {
	e, err := getCollyHTMLElement()
	assert.NoError(t, err, "getCollyHTMLElement failed")

	match, err := getMatchResult(&e)

	assert.NoError(t, err, "getMatchResult failed")

	assert.NotNil(t, match, "getMatchResult returned nil")
	assert.NotNil(t, match.URL, "URL is nil")
	assert.NotNil(t, match.Team1, "Team1 is nil")
	assert.NotNil(t, match.Team2, "Team2 is nil")
	assert.NotNil(t, match.Score, "Score is nil")
	assert.NotNil(t, match.StartTime, "StartTime is nil")
	assert.NotNil(t, match.Tag, "Tag is nil")
	assert.NotNil(t, match.Status, "Status is nil")
	assert.NotNil(t, match.Region, "Region is nil")

	assert.Equal(t, match.Status, Completed, "Status is not Completed")
}

func TestFindRegion(t *testing.T) {
	testTags := map[string]string{
		"Challengers League 2024 Spain Rising: Split 2":    "ES",
		"Challengers League 2024 Americas Rising: Split 2": "NA",
		"Challengers League 2024 EMEA Rising: Split 2":     "EU",
		"Challengers League 2024 Korea Rising: Split 2":    "KR",
		"Challengers League 2024 Brazil Rising: Split 2":   "BR",
		"Challengers League 2024 Japan Rising: Split 2":    "JP",
		"Challengers League 2024 LATAM Rising: Split 2":    "LATAM",
	}

	for tag, region := range testTags {
		assert.Equal(t, region, findRegion(tag), "Region mismatch")
	}
}

func getCollyHTMLElement() (colly.HTMLElement, error) {
	collyElement := colly.HTMLElement{}

	c := colly.NewCollector()
	c.OnHTML("div.js-home-matches-upcoming", func(e *colly.HTMLElement) {
		e.ForEach("a.wf-module-item", func(_ int, el *colly.HTMLElement) {
			collyElement = *el
		})
	})

	c.Visit("https://www.vlr.gg")
	c.Wait()

	return collyElement, nil
}
