package scrapper

import (
	"reflect"
	"testing"
)

func TestGetMatches(t *testing.T) {
	expected := []Match{} // replace this with the expected result
	result := GetMatches()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMatches() = %v; want %v", result, expected)
	}
}
