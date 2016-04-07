package players

import (
	"github.com/skiesel/joshuatree/domains"
)

type Player interface {
	GetAction(domain domains.Domain, state domains.State, opposingActions []domains.Action) domains.Action
}
