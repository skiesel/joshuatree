package main

import (
	"github.com/skiesel/mcts/domains"
	"github.com/skiesel/mcts/games"
	"github.com/skiesel/mcts/players"
)

func main() {
	tictactoe := domains.TicTacToe{}
	players := []players.Player{
		players.NewMonteCarloPlayer(0, []int64{1}, 5., 0.5, tictactoe),
		players.NewInteractivePlayer(),
	}
	games.Play(tictactoe, players)
}
