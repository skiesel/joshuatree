package players

import (
	"github.com/skiesel/mcts/domains"
)

type Player interface {
	GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action
}
