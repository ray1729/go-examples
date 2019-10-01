package tictactoe

import (
	"math/rand"

	"github.com/google/uuid"
)

type RandomPlayer struct {
	id uuid.UUID
}

func NewRandomPlayer() Player {
	return RandomPlayer{id: uuid.New()}
}

func (p RandomPlayer) Id() uuid.UUID {
	return p.id
}

func (p RandomPlayer) NextMove(g GameState) uint8 {
	avail := g.AvailableMoves()
	n := rand.Intn(len(avail))
	return avail[n]
}
