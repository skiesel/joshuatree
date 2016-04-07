package players

import (
	"github.com/skiesel/joshuatree/domains"
	"github.com/skiesel/joshuatree/policies/exterior"
)

type RandomPlayer struct {
	policy *exterior.RandomPolicy
}

func NewRandomPlayer() *RandomPlayer {
	return &RandomPlayer{
		policy: exterior.NewRandomPolicy(),
	}
}

func (player RandomPlayer) GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action {
	return player.policy.GetAction(domain, state)
}
