package main

// Scores implements Heap interface
type Scores []Score

func (scores Scores) Less(i, j int) bool {
	// MaxHear - invert order
	return scores[i].GetTotalScore() >= scores[j].GetTotalScore()
}

func (scores Scores) Len() int {
	return len(scores)
}

func (scores Scores) Swap(i, j int) {
	scores[i], scores[j] = scores[j], scores[i]
}

func (scores *Scores) Push(x interface{}) {
	*scores = append(*scores, x.(Score))
}

func (scores *Scores) Pop() interface{} {
	old := *scores
	n := len(old)
	x := old[n-1]
	*scores = old[0 : n-1]
	return x
}
