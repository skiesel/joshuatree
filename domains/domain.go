package domains

const (
	TIE = -1
)

type Domain interface {
	GetAvailableActions(state State) []Action
	ApplyAction(state State, action Action, playerIndex int64) State

	GetStartState() State

	IsTerminal(state State) bool
	DidWin(state State, playerIndex int64) bool
	WhoWon(state State) int64

	Draw(state State)
	HashState(state State) int64
	CompareStates(state1, state2 State) int64

	CompareActions(action1, action2 Action) int64
	GetString(action Action) string
	HashAction(action Action) int64
}

type State interface {
}

type Action interface {
}
