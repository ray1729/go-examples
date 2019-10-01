package tictactoe

import "fmt"

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

type Player struct {
	Name string
}

type GameState struct {
	Players []Player
	Moves   []uint8
	Winner  *int
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

func (g GameState) Play(move uint8) (GameState, error) {
	if move < 0 || move > 8 || Bitmask(g.Moves, 0, 1)&(1<<move) != 0 {
		return g, fmt.Errorf("Illegal move: %d", move)
	}
	if g.Winner != nil {
		return g, fmt.Errorf("Game over, player %d won", *g.Winner)
	}
	player := len(g.Moves) % 2
	g.Moves = append(g.Moves, move)
	if g.HasWon(player) {
		g.Winner = &player
	}
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
