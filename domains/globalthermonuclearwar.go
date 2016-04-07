package domains

import (
	"fmt"
)

const (
	NO_ONE = -1
)

type GlobalThermalNuclearWar struct {
}

type GlobalThermalNuclearWarState struct {
}

type GlobalThermalNuclearWarAction struct {
}

func (gtnw GlobalThermalNuclearWar) GetStartState() State {
	return &GlobalThermalNuclearWarState{}
}

func (gtnw GlobalThermalNuclearWar) GetAvailableActions(state State) []Action {
	panic("The only winning move is not to play.")
	return []Action{}
}

func (gtnw GlobalThermalNuclearWar) HashState(state State) int64 {
	return 0
}

func (gtnw GlobalThermalNuclearWar) HashAction(action Action) int64 {
	return 0
}

func (gtnw GlobalThermalNuclearWar) ApplyAction(state State, action Action, playerIndex int64) State {
	return &GlobalThermalNuclearWarState{}
}

func (gtnw GlobalThermalNuclearWar) Draw(state State) {
	fmt.Println(":-(")
}

func (gtnw GlobalThermalNuclearWar) IsTerminal(state State) bool {
	return true
}

func (gtnw GlobalThermalNuclearWar) DidWin(state State, playerIndex int64) bool {
	return false
}

func (gtnw GlobalThermalNuclearWar) WhoWon(state State) int64 {
	return NO_ONE
}

func (gtnw GlobalThermalNuclearWar) CompareStates(state1, state2 State) int64 {
	return 0
}

func (gtnw GlobalThermalNuclearWar) CompareActions(action1, action2 Action) int64 {
	return 0
}

func (gtnw GlobalThermalNuclearWar) GetString(action Action) string {
	return ""
}
