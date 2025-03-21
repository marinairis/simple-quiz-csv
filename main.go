package main

import (
	"github.com/marinairis/quiz-go/game"
)

func main() {
	game := &game.GameState{Points: 0}

	game.Init()
	themeFile := game.ChooseTheme()
	game.ProccessCSV(themeFile)

	game.Run()
	game.CheckApproval()
}
