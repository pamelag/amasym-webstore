package util

import "testing"

func TestRound(t *testing.T) {
	
	expectedResult := 29.78
	
	actualResult := Round(29.778)
	
	if actualResult != expectedResult {
		t.Fatalf("Expected %f but got %f", expectedResult, actualResult)
	}
}
