package scrapper

import (
	"testing"
)

func TestGetMatchesAndResults(t *testing.T) {
	matchesResult := GetMatchesFromVLR()
	resultsResult := GetResultsFromVLR()

	if len(matchesResult) == 0 {
		t.Error("No matches found")
	}

	if len(resultsResult) == 0 {
		t.Error("No results found")
	}
}
