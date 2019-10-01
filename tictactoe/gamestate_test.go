package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameState(t *testing.T) {
	moves := []uint8{0, 3, 1, 4, 2}
	g := GameState{}
	var err error
	for _, m := range moves {
		g, err = g.Play(m)
		assert.Nil(t, err)
	}
	assert.Equal(t, 0, *g.Winner, "Player 0 wins")
}

func TestAvailableMoves(t *testing.T) {
	g := GameState{}
	assert.Equal(t, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}, g.AvailableMoves(), "All moves available from starting position")
	for move := uint8(0); move < 9; move++ {
		g2, err := g.Play(move)
		assert.Nil(t, err, "%d is a valid move", move)
		avail := g2.AvailableMoves()
		assert.Equal(t, 8, len(avail), "There are 8 moves available")
		assert.NotContains(t, avail, move, "Move %d is not available", move)
	}
}
