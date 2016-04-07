package exterior_policies

import (
	"github.com/skiesel/mcts/domains"
	"math/rand"
)

type ExteriorTreePolicy interface {
	GetAction(domain domains.Domain, state domains.State) domains.Action
}

type RandomPolicy struct{}

func NewRandomPolicy() *RandomPolicy {
	return &RandomPolicy{}
}

func (policy RandomPolicy) GetAction(domain domains.Domain, state domains.State) domains.Action {
	actions := domain.GetAvailableActions(state)
	return actions[rand.Intn(len(actions))]
}
