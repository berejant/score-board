package main

import (
	"os"
	"syscall/js"
)

import (
	"encoding/json"
	"os/signal"
	"syscall"
)

var board = ScoreBoard{}

func main() {
	js.Global().Set("PutScore", js.FuncOf(PutScore))
	js.Global().Set("GetScores", js.FuncOf(GetScores))

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until user hits ctrl+c
}

func PutScore(this js.Value, args []js.Value) interface{} {
	score, err := MakeScoreFromString(args[0].String())

	if err == nil {
		board.AddScore(score)
		return ""
	} else {
		return js.ValueOf(err.Error())
	}
}

func GetScores(this js.Value, args []js.Value) interface{} {
	scores := board.GetScores()
	jsonScores, _ := json.Marshal(scores)

	return js.ValueOf(string(jsonScores))
}
