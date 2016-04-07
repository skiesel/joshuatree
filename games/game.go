package games

import (
	"fmt"
	"github.com/skiesel/joshuatree/domains"
	"github.com/skiesel/joshuatree/players"
	"math/rand"
	"time"
)

func Play(domain domains.Domain, players []players.Player) {
	rand.Seed(time.Now().Unix())

	numPlayers := len(players)
	turns := getRandomPlayOrder(numPlayers)

	turn := 0
	currentState := domain.GetStartState()

	domain.Draw(currentState)

	actionHistory := []domains.Action{}

	for !domain.IsTerminal(currentState) {
		player := turns[turn]
		turn++
		if turn >= len(turns) {
			turn = 0
		}

		action := players[player].GetAction(domain, currentState, actionHistory)
		currentState = domain.ApplyAction(currentState, action, player)

		actionHistory = append(actionHistory, action)
		if len(actionHistory) > len(players) {
			actionHistory = actionHistory[1:]
		}

		domain.Draw(currentState)
	}

	whoWon := domain.WhoWon(currentState)
	if whoWon == domains.TIE {
		fmt.Printf("The game was a tie!\n")
	} else {
		fmt.Printf("Player %d won!\n", whoWon)
	}
}

func getRandomPlayOrder(numPlayers int) []int64 {
	turns := make([]int64, numPlayers)
	for i := range turns {
		turns[i] = int64(i)
	}

	numShuffles := numPlayers * 10
	for i := 0; i < numShuffles; i++ {
		index1 := rand.Intn(numPlayers)
		index2 := rand.Intn(numPlayers)
		temp := turns[index1]
		turns[index1] = turns[index2]
		turns[index2] = temp
	}

	return turns
}
