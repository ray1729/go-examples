package tictactoe

import "fmt"

func GameLoop(player1, player2 Player) GameState {
	game := GameState{}
	players := [2]Player{player1, player2}
	for !game.IsGameOver() {
		player := players[len(game.Moves) % 2]
		for {
			move := player.NextMove(game)
			nextState, err := game.Play(move)
			if err != nil {
				fmt.Println(err)
				continue
			}
			game = nextState
			break
		}
	}
	return game
}
