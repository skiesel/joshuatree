package exterior

import (
	"github.com/skiesel/joshuatree/domains"
	"math/rand"
)

type RandomPolicy struct{}

func NewRandomPolicy() *RandomPolicy {
	return &RandomPolicy{}
}

func (policy RandomPolicy) GetAction(domain domains.Domain, state domains.State) domains.Action {
	actions := domain.GetAvailableActions(state)
	return actions[rand.Intn(len(actions))]
}
