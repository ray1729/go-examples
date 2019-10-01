package tictactoe

import (
	"fmt"

	"github.com/google/uuid"
)

// Assuming a 3x3 grid like so:
// 0 | 1 | 2
// ---+---+---
// 3 | 4 | 5
// ---+---+--
// 6 | 7 | 8
//
// ...this var represents the winning positions

var WinningPositions = [8][]uint8{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

// We represent these same positions as bitmasks

var WinningPositionMasks [8]uint16

func init() {
	for i, p := range WinningPositions {
		WinningPositionMasks[i] = Bitmask(p, 0, 1)
	}
}

type Player interface {
	Id() uuid.UUID
	NextMove(g GameState) uint8
}

type GameState struct {
	Moves []uint8
}

func Bitmask(xs []uint8, start, step int) uint16 {
	var mask uint16
	for i := start; i < len(xs); i += step {
		mask |= (1 << xs[i])
	}
	return mask
}

func (g GameState) HasWon(player int) bool {
	position := Bitmask(g.Moves, player, 2)
	for _, winningPosition := range WinningPositionMasks {
		if position&winningPosition == winningPosition {
			return true
		}
	}
	return false
}

func (g GameState) IsGameOver() bool {
	return len(g.Moves) == 9 || g.HasWon(0) || g.HasWon(1)
}

func (g GameState) Play(move uint8) (GameState, error) {
	if move < 0 || move > 8 || Bitmask(g.Moves, 0, 1)&(1<<move) != 0 {
		return g, fmt.Errorf("Illegal move: %d", move)
	}
	if g.IsGameOver() {
		return g, fmt.Errorf("This game is over")
	}
	g.Moves = append(g.Moves, move)
	return g, nil
}

func (g GameState) AvailableMoves() []uint8 {
	var moves []uint8
	mask := Bitmask(g.Moves, 0, 1)
	for move := uint8(0); move < 9; move++ {
		if mask&(1<<move) == 0 {
			moves = append(moves, move)
		}
	}
	return moves
}

func PlayerSymbol(player int) string {
	if player % 2 == 0 {
		return "X"
	}
	return "O"
}

func (g GameState) FormatGrid() string {
	xs := []interface{}{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	for i, p := range g.Moves {
		xs[p] = PlayerSymbol(i)
	}
	return fmt.Sprintf(" %s | %s | %s\n---+---+---\n %s | %s | %s\n---+---+---\n %s | %s | %s\n", xs...)
}

func (g GameState) Status() string {
	for p := 0; p < 2; p++ {
		if g.HasWon(p) {
			return fmt.Sprintf("%s's won", PlayerSymbol(p))
		}
	}
	if len(g.Moves) == 9 {
		return "It's a draw"
	}
	return fmt.Sprintf("%s to play", PlayerSymbol(len(g.Moves)))
}

func (g GameState) String() string {
	return g.FormatGrid() + "\n" + g.Status() + "\n"
}
