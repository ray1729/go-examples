package tictactoe

import "fmt"

func GameLoop(player1, player2 Player) GameState {
	state := GameState{}
	players := [2]Player{player1, player2}
	for !state.IsGameOver() {
		player := players[len(state.Moves) % 2]
		move := player.NextMove(state)
		nextState, err := state.Play(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		state = nextState
	}
	return state
}
