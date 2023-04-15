package main

import (
	"reflect"
	"testing"
)

func TestMakeScoreFromString(t *testing.T) {
	t.Run("success 1", func(t *testing.T) {
		expectedScore := Score{
			HomeTeamName:   "Mexico",
			HomeTeamScore:  0,
			GuestTeamName:  "Canada",
			GuestTeamScore: 5,
		}

		actualScore, err := MakeScoreFromString("Mexico - Canada: 0 - 5")

		if !reflect.DeepEqual(expectedScore, actualScore) {
			t.Errorf("want: %v; got: %v", expectedScore, actualScore)
		}

		if err != nil {
			t.Errorf("want: not error; got: %v", err)
		}
	})

	t.Run("success 2", func(t *testing.T) {
		expectedScore := Score{
			HomeTeamName:   "Germany",
			HomeTeamScore:  2,
			GuestTeamName:  "France",
			GuestTeamScore: 2,
		}

		actualScore, err := MakeScoreFromString("Germany - France: 2 – 2")

		if !reflect.DeepEqual(expectedScore, actualScore) {
			t.Errorf("want: %v; got: %v", expectedScore, actualScore)
		}

		if err != nil {
			t.Errorf("want: not error; got: %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		expectedScore := Score{
			HomeTeamName:   "",
			HomeTeamScore:  0,
			GuestTeamName:  "",
			GuestTeamScore: 0,
		}

		actualScore, err := MakeScoreFromString("Mexico - Canada: LA - 5")

		if !reflect.DeepEqual(expectedScore, actualScore) {
			t.Errorf("want: %v; got: %v", expectedScore, actualScore)
		}

		if err == nil {
			t.Errorf("want: error; got: nil")
		}
	})
}
