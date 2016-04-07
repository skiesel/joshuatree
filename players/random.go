package players

import (
	"github.com/skiesel/mcts/domains"
	"github.com/skiesel/mcts/policies/exterior"
)

type RandomPlayer struct {
	policy *exterior_policies.RandomPolicy
}

func NewRandomPlayer() *RandomPlayer {
	return &RandomPlayer{
		policy: exterior_policies.NewRandomPolicy(),
	}
}

func (player RandomPlayer) GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action {
	return player.policy.GetAction(domain, state)
}
