package scrapper

import (
	"net/http"
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
)

func TestGetMatch(t *testing.T) {
	e := getElementFromFile(t, "/test_data/matches.html")
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
	e := getElementFromFile(t, "test_data/results.html")
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
		"Challengers League 2024 Spain Rising: Split 2":    "es",
		"Challengers League 2024 Americas Rising: Split 2": "na",
		"Challengers League 2024 EMEA Rising: Split 2":     "emea",
		"Challengers League 2024 Korea Rising: Split 2":    "kr",
		"Challengers League 2024 Brazil Rising: Split 2":   "br",
		"Challengers League 2024 Japan Rising: Split 2":    "jp",
		"Challengers League 2024 LATAM Rising: Split 2":    "latam",
	}

	for tag, region := range testTags {
		assert.Equal(t, region, findRegion(tag), "Region mismatch")
	}
}

func getElementFromFile(t *testing.T, path string) colly.HTMLElement {
	tport := &http.Transport{}
	tport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(tport)

	var collyElement colly.HTMLElement
	c.OnHTML("div.wf-module-item", func(e *colly.HTMLElement) {
		collyElement = *e
	})

	c.Visit("file://" + path)

	c.Wait()

	c.OnError(func(r *colly.Response, err error) {
		t.Fatalf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	return collyElement
}
