package game

type GameState struct {
	Name      string
	Points    int
	Questions []Question
	Approved  bool
}
