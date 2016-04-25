package domains

import (
	"fmt"
)

type TicTacToe struct {
}

type TicTacToeState struct {
	board []int64
}

type TicTacToeAction struct {
	which int64
}

func (tictactoe TicTacToe) GetStartState() State {
	board := make([]int64, 9)
	for i := 0; i < 9; i++ {
		board[i] = TIE
	}

	return TicTacToeState{
		board: board,
	}
}

func (tictactoe TicTacToe) GetAvailableActions(state State) []Action {
	actions := []Action{}
	tictactoeState := state.(TicTacToeState)
	for i := range tictactoeState.board {
		if tictactoeState.board[i] == TIE {
			actions = append(actions, TicTacToeAction{which: int64(i)})
		}
	}
	return actions
}

func (tictactoe TicTacToe) StateString(state State) string {
	tictactoeState := state.(TicTacToeState)
	str := ""
	for _, cell := range tictactoeState.board {
		if cell == TIE {
			str += "3"
		} else {
			str += fmt.Sprintf("%d", cell)
		}
	}
	return str
}

func (tictactoe TicTacToe) ActionString(action Action) string {
	tictactoeAction := action.(TicTacToeAction)
	return fmt.Sprintf("%d", tictactoeAction.which)
}

func (tictactoe TicTacToe) ApplyAction(state State, action Action, playerIndex int64) State {
	tictactoeState := state.(TicTacToeState)
	tictactoeAction := action.(TicTacToeAction)

	nextState := TicTacToeState{
		board: make([]int64, 9),
	}

	for i := range tictactoeState.board {
		nextState.board[i] = tictactoeState.board[i]
	}

	nextState.board[tictactoeAction.which] = playerIndex

	return nextState
}

func (tictactoe TicTacToe) Draw(state State) {
	tictactoeState := state.(TicTacToeState)
	for i, board := range tictactoeState.board {
		if i > 0 && i%3 == 0 {
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

func (tictactoe TicTacToe) IsTerminal(state State) bool {
	noMoreMoves := len(tictactoe.GetAvailableActions(state)) == 0
	return noMoreMoves || tictactoe.WhoWon(state) != TIE
}

func (tictactoe TicTacToe) DidWin(state State, playerIndex int64) bool {
	return tictactoe.WhoWon(state) == playerIndex
}

func (tictactoe TicTacToe) WhoWon(state State) int64 {
	tictactoeState := state.(TicTacToeState)
	horizontalWin := checkHorizontalWinTicTacToe(tictactoeState)
	verticalWin := checkVerticalWinTicTacToe(tictactoeState)
	diagonalWin := checkDiagonalWinTicTacToe(tictactoeState)
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

func checkHorizontalWinTicTacToe(state TicTacToeState) int64 {
	if state.board[0] != TIE && state.board[0] == state.board[1] && state.board[1] == state.board[2] {
		return state.board[0]
	}
	if state.board[3] != TIE && state.board[3] == state.board[4] && state.board[4] == state.board[5] {
		return state.board[3]
	}
	if state.board[6] != TIE && state.board[6] == state.board[7] && state.board[7] == state.board[8] {
		return state.board[6]
	}
	return TIE
}

func checkVerticalWinTicTacToe(state TicTacToeState) int64 {
	if state.board[0] != TIE && state.board[0] == state.board[3] && state.board[3] == state.board[6] {
		return state.board[0]
	}
	if state.board[1] != TIE && state.board[1] == state.board[4] && state.board[4] == state.board[7] {
		return state.board[1]
	}
	if state.board[2] != TIE && state.board[2] == state.board[5] && state.board[5] == state.board[8] {
		return state.board[2]
	}
	return TIE
}

func checkDiagonalWinTicTacToe(state TicTacToeState) int64 {
	if state.board[0] != TIE && state.board[0] == state.board[4] && state.board[4] == state.board[8] {
		return state.board[0]
	}
	if state.board[2] != TIE && state.board[2] == state.board[4] && state.board[4] == state.board[6] {
		return state.board[2]
	}
	return TIE
}

func (tictactoe TicTacToe) CompareStates(state1, state2 State) int64 {
	tictactoeState1 := state1.(TicTacToeState)
	tictactoeState2 := state2.(TicTacToeState)
	for i := range tictactoeState1.board {
		if tictactoeState1.board[i] != tictactoeState2.board[i] {
			return -1
		}
	}
	return 0
}

func (tictactoe TicTacToe) CompareActions(action1, action2 Action) int64 {
	tictactoeAction1 := action1.(TicTacToeAction)
	tictactoeAction2 := action2.(TicTacToeAction)
	if tictactoeAction1.which != tictactoeAction2.which {
		return -1
	}
	return 0
}

func (tictactoe TicTacToe) GetString(action Action) string {
	tictactoeAction := action.(TicTacToeAction)
	return fmt.Sprintf("%d", tictactoeAction.which)
}
