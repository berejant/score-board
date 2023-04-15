package main

import (
	"reflect"
	"testing"
)

func TestScoreBoard_GetScores(t *testing.T) {
	expectedScores := []Score{
		{
			HomeTeamName:   "Uruguay",
			HomeTeamScore:  6,
			GuestTeamName:  "Italy",
			GuestTeamScore: 6,
		},
		{
			HomeTeamName:   "Spain",
			HomeTeamScore:  10,
			GuestTeamName:  "Brazil",
			GuestTeamScore: 2,
		},
		{
			HomeTeamName:   "Mexico",
			HomeTeamScore:  0,
			GuestTeamName:  "Canada",
			GuestTeamScore: 5,
		},
		{
			HomeTeamName:   "Argentina",
			HomeTeamScore:  3,
			GuestTeamName:  "Australia",
			GuestTeamScore: 1,
		},
		{
			HomeTeamName:   "Germany",
			HomeTeamScore:  2,
			GuestTeamName:  "France",
			GuestTeamScore: 2,
		},
	}

	board := ScoreBoard{}

	board.AddScore(expectedScores[2])
	board.AddScore(expectedScores[1])
	board.AddScore(expectedScores[4])
	board.AddScore(expectedScores[0])
	board.AddScore(expectedScores[3])

	boardScores := board.GetScores()

	if !reflect.DeepEqual(expectedScores, boardScores) {
		t.Errorf("want: %v; got: %v", expectedScores, boardScores)
	}
}
