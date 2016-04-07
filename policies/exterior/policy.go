package exterior

import (
	"github.com/skiesel/joshuatree/domains"
)

type ExteriorTreePolicy interface {
	GetAction(domain domains.Domain, state domains.State) domains.Action
}