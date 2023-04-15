package main

import (
	"container/heap"
)

type ScoreBoard struct {
	scores Scores
}

func (board *ScoreBoard) AddScore(score Score) {
	heap.Push(&board.scores, score)
}

func (board *ScoreBoard) GetScores() []Score {
	sortedScores := Scores{}
	_scores := make(Scores, board.scores.Len())
	copy(_scores, board.scores)

	for _scores.Len() > 0 {
		sortedScores = append(sortedScores, heap.Pop(&_scores).(Score))
	}

	return sortedScores
}
