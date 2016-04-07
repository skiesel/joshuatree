package players

import (
	"github.com/skiesel/mcts/domains"
	"github.com/skiesel/mcts/search"
)

type MonteCarloPlayer struct {
	monteCarlo *search.MonteCarlo
}

func NewMonteCarloPlayer(playerId int64, otherPlayerIds []int64, moveTimeout, uctCValue float64, domain domains.Domain) *MonteCarloPlayer {
	return &MonteCarloPlayer{
		monteCarlo: search.NewMonteCarlo(playerId, otherPlayerIds, moveTimeout, uctCValue, domain),
	}
}

func (player MonteCarloPlayer) GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action {
	player.monteCarlo.UpdateForOpposingAction(domain, state, opposingActions)
	return player.monteCarlo.GetAction(domain)
}
