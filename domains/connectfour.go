package domains

import (
	"fmt"
)

type ConnectFour struct {
}

type ConnectFourState struct {
	board []int64
}

type ConnectFourAction struct {
	which int64
}

func (connectfour ConnectFour) GetStartState() State {
	board := make([]int64, 42)
	for i := 0; i < 42; i++ {
		board[i] = TIE
	}

	return ConnectFourState{
		board: board,
	}
}

func (connectfour ConnectFour) GetAvailableActions(state State) []Action {
	actions := []Action{}
	connectfourState := state.(ConnectFourState)
	for i := 0; i < 7; i++ {
		if connectfourState.board[35+i] == TIE {
			actions = append(actions, ConnectFourAction{which: int64(i)})
		}
	}
	return actions
}

func (connectfour ConnectFour) StateString(state State) string {
	connectfourState := state.(ConnectFourState)
	str := ""
	for _, cell := range connectfourState.board {
		if cell == TIE {
			str += "3"
		} else {
			str += fmt.Sprintf("%d", cell)
		}
	}
	return str
}

func (connectfour ConnectFour) ActionString(action Action) string {
	connectfourAction := action.(ConnectFourAction)
	return fmt.Sprintf("%d", connectfourAction.which)
}

func (connectfour ConnectFour) ApplyAction(state State, action Action, playerIndex int64) State {
	connectfourState := state.(ConnectFourState)
	connectfourAction := action.(ConnectFourAction)

	nextState := ConnectFourState{
		board: make([]int64, 42),
	}

	positionSet := false
	for i := range connectfourState.board {
		if !positionSet && int64(i%7) == connectfourAction.which && connectfourState.board[i] == TIE {
			nextState.board[i] = playerIndex
			positionSet = true
		} else {
			nextState.board[i] = connectfourState.board[i]
		}
	}

	return nextState
}

func (connectfour ConnectFour) Draw(state State) {
	connectfourState := state.(ConnectFourState)
	for i, board := range connectfourState.board {
		if i > 0 && i%7 == 0 {
			fmt.Print("\n")
		}
		if board == TIE {
			fmt.Print("_")
		} else {
			fmt.Print(board)
		}
	}
	fmt.Print("\n\n")
}

func (connectfour ConnectFour) IsTerminal(state State) bool {
	connectfourState := state.(ConnectFourState)
	noMoreMoves := len(connectfour.GetAvailableActions(state)) == 0
	return noMoreMoves || checkHorizontalWinConnectFour(connectfourState) != TIE || checkVerticalWinConnectFour(connectfourState) != TIE || checkDiagonalWinConnectFour(connectfourState) != TIE
}

func (connectfour ConnectFour) DidWin(state State, playerIndex int64) bool {
	connectfourState := state.(ConnectFourState)
	return checkHorizontalWinConnectFour(connectfourState) == playerIndex || checkVerticalWinConnectFour(connectfourState) == playerIndex || checkDiagonalWinConnectFour(connectfourState) == playerIndex
}

func (connectfour ConnectFour) WhoWon(state State) int64 {
	connectfourState := state.(ConnectFourState)
	horizontalWin := checkHorizontalWinConnectFour(connectfourState)
	verticalWin := checkVerticalWinConnectFour(connectfourState)
	diagonalWin := checkDiagonalWinConnectFour(connectfourState)
	if horizontalWin != TIE {
		return horizontalWin
	}
	if verticalWin != TIE {
		return verticalWin
	}
	if diagonalWin != TIE {
		return diagonalWin
	}
	return TIE
}

func checkHorizontalWinConnectFour(state ConnectFourState) int64 {
	for i := 0; i < 6; i++ { //row
		for j := 0; j < 4; j++ { //column
			winner := TIE
			for k := 0; k < 4; k++ { //column
				index := i*7 + j + k
				if state.board[index] == TIE {
					winner = TIE
					break
				} else if winner == TIE {
					winner = state.board[index]
				} else if winner != state.board[index] {
					winner = TIE
					break
				}
			}
			if winner != TIE {
				return winner
			}
		}
	}
	return TIE
}

func checkVerticalWinConnectFour(state ConnectFourState) int64 {
	for i := 0; i < 7; i++ { //column
		for j := 0; j < 3; j++ { //row
			winner := TIE
			for k := 0; k < 4; k++ { //row
				index := 7*(j+k) + i
				if state.board[index] == TIE {
					winner = TIE
					break
				} else if winner == TIE {
					winner = state.board[index]
				} else if winner != state.board[index] {
					winner = TIE
					break
				}
			}
			if winner != TIE {
				return winner
			}
		}
	}
	return TIE
}

func checkDiagonalWinConnectFour(state ConnectFourState) int64 {
	for i := 0; i < 3; i++ { //row
		for j := 0; j < 4; j++ { //column
			winner := TIE
			for k := 0; k < 4; k++ { //column
				index := (i+k)*7 + j + k
				if state.board[index] == TIE {
					winner = TIE
					break
				} else if winner == TIE {
					winner = state.board[index]
				} else if winner != state.board[index] {
					winner = TIE
					break
				}
			}
			if winner != TIE {
				return winner
			}
		}
	}

	for i := 5; i >= 3; i-- { //row
		for j := 0; j < 4; j++ { //column
			winner := TIE
			for k := 0; k < 4; k++ { //column
				index := (i-k)*7 + j + k
				if state.board[index] == TIE {
					winner = TIE
					break
				} else if winner == TIE {
					winner = state.board[index]
				} else if winner != state.board[index] {
					winner = TIE
					break
				}
			}
			if winner != TIE {
				return winner
			}
		}
	}

	return TIE
}

func (connectfour ConnectFour) CompareStates(state1, state2 State) int64 {
	connectfourState1 := state1.(ConnectFourState)
	connectfourState2 := state2.(ConnectFourState)
	for i := range connectfourState1.board {
		if connectfourState1.board[i] != connectfourState2.board[i] {
			return -1
		}
	}
	return 0
}

func (connectfour ConnectFour) CompareActions(action1, action2 Action) int64 {
	connectfourAction1 := action1.(ConnectFourAction)
	connectfourAction2 := action2.(ConnectFourAction)
	if connectfourAction1.which != connectfourAction2.which {
		return -1
	}
	return 0
}

func (connectfour ConnectFour) GetString(action Action) string {
	connectfourAction := action.(ConnectFourAction)
	return fmt.Sprintf("%d", connectfourAction.which)
}
