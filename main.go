package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/skiesel/joshuatree/domains"
	"github.com/skiesel/joshuatree/games"
	"github.com/skiesel/joshuatree/players"
)

func playTicTacToe() {
	tictactoe := domains.TicTacToe{}
	players := []players.Player{
		players.NewMonteCarloPlayer(0, []int64{1}, 1., 0.5, tictactoe),
		players.NewInteractivePlayer(),
	}
	games.Play(tictactoe, players)
}

func playConnectFour() {
	connectfour := domains.ConnectFour{}
	players := []players.Player{
		players.NewMonteCarloPlayer(0, []int64{1}, 1., 0.5, connectfour),
		players.NewInteractivePlayer(),
	}
	games.Play(connectfour, players)
}

func playGlobalThermalNuclearWar() {
	gtnw := domains.GlobalThermalNuclearWar{}
	players := []players.Player{
		players.NewMonteCarloPlayer(0, []int64{1}, 1., 0.5, gtnw),
		players.NewInteractivePlayer(),
	}
	games.Play(gtnw, players)
}

func main() {
	type GamePair struct {
		Label string
		Play  func()
	}

	availableGames := []GamePair{GamePair{Label: "TicTacToe", Play: playTicTacToe},
		GamePair{Label: "ConnectFour", Play: playConnectFour},
		GamePair{Label: "GlobalThermalNuclearWar", Play: playGlobalThermalNuclearWar},
		GamePair{Label: "Quit", Play: func() {}},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please choose a game:")
		for i, pair := range availableGames {
			fmt.Printf("%d) %s\n", i+1, pair.Label)
		}
		fmt.Printf(">")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		selection, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("failed to parse selection")
			continue
		}
		selection--
		if selection < 0 || selection >= int64(len(availableGames)) {
			fmt.Println("selection is not a valid game number")
			continue
		}

		if selection == int64(len(availableGames)-1) {
			break
		}

		availableGames[selection].Play()
	}
}
