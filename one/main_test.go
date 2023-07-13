package main

import "testing"

const (
	CALORIES = 24000
)

func TestGetMaxCalories(t *testing.T) {
	calories, _ := getMaxCalories("test.txt")

	if calories != CALORIES {
		t.Errorf("expected calories %d; got %d", CALORIES, calories)
	}
}