package main

import (
	"errors"
	"regexp"
	"strconv"
)

type Score struct {
	HomeTeamName  string // should be separate entity Team { Id int, Name string ... }
	HomeTeamScore int

	GuestTeamName  string
	GuestTeamScore int
}

func (score Score) GetTotalScore() int {
	return score.HomeTeamScore + score.GuestTeamScore
}

// https://regex101.com/r/v5DthJ/1
var scoreStringRegExp = regexp.MustCompile(`^\s*(\w+)\s*\W\s*(\w+)\s*\W\s*(\d+)\s*\W\s*(\d+)\w*$`)

func MakeScoreFromString(input string) (score Score, err error) {
	matched := scoreStringRegExp.FindStringSubmatch(input)

	if len(matched) == 0 {
		err = errors.New(`Incorrect input score string: ` + input)
	} else {
		score.HomeTeamName = matched[1]
		score.GuestTeamName = matched[2]
		score.HomeTeamScore, _ = strconv.Atoi(matched[3])
		score.GuestTeamScore, _ = strconv.Atoi(matched[4])
	}
	return
}
