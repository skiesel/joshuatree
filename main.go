package main

import (
	"github.com/skiesel/joshuatree/domains"
	"github.com/skiesel/joshuatree/games"
	"github.com/skiesel/joshuatree/players"
)

func main() {
	tictactoe := domains.TicTacToe{}
	players := []players.Player{
		players.NewMonteCarloPlayer(0, []int64{1}, 5., 0.5, tictactoe),
		players.NewInteractivePlayer(),
	}
	games.Play(tictactoe, players)
}
