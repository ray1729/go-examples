package main

import (
	"fmt"

	"github.com/ray1729/go-examples/tictactoe"
)

func main() {
	player1 := tictactoe.NewRandomPlayer()
	player2 := tictactoe.NewRandomPlayer()
	var p1win, p2win, draw int
	for i := 0; i < 5000; i++ {
		game := tictactoe.GameLoop(player1, player2)
		if game.HasWon(0) {
			p1win++
		} else if game.HasWon(1) {
			p2win++
		} else {
			draw++
		}
	}
	fmt.Printf("Player 1: %d, Player 2: %d, Draw: %d\n", p1win, p2win, draw)
}
