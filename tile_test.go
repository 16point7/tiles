package tiles

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()

	if g.score != 0 {
		t.Fatalf("New game must have score of 0. got %d", g.score)
	}

	if len(g.board) != 4 {
		t.Fatalf("Invalid number of rows. got %d, want %d", len(g.board), length)
	}

	for j, row := range g.board {
		if len(row) != length {
			t.Fatalf("Row %d has invalid number of columns. got %d, want %d", j, len(row), length)
		}
	}

	var initVals []uint
	for _, row := range g.board {
		for _, val := range row {
			if val != 0 {
				initVals = append(initVals, val)
			}
		}
	}

	if len(initVals) != numInitVals {
		t.Fatalf("Wrong number of initial values. got %d, want %d", len(initVals), numInitVals)
	}

	for _, val := range initVals {
		if val != 2 && val != 4 {
			t.Fatalf("Invalid initial value. got %d, want 2 or 4", val)
		}
	}
}

type moveRes struct {
	want     [][]uint
	score    uint
	gameOver bool
	newTile  bool
}

func TestMove(t *testing.T) {
	tests := []struct {
		board [][]uint
		down  moveRes
		up    moveRes
	}{
		{
			board: [][]uint{
				{2, 4, 0, 0},
				{2, 4, 0, 8},
				{2, 4, 0, 0},
				{2, 4, 0, 0},
			},
			down: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{4, 8, 0, 0},
					{4, 8, 0, 8},
				},
				score:    24,
				gameOver: false,
				newTile:  true,
			},
			up: moveRes{
				want: [][]uint{
					{4, 8, 0, 8},
					{4, 8, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				score:    24,
				gameOver: false,
				newTile:  true,
			},
		},
	}

	g := &game{}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			g.board = test.board

			gameOver := g.MoveDown()
			assertMoveRes(t, g.board, test.down.want, gameOver, test.down.gameOver, test.down.newTile, g.score, test.down.score)

		})
	}
}

func assertMoveRes(t *testing.T, gotBoard, wantBoard [][]uint, gotGameOver, wantGameOver, newTile bool, gotScore, wantScore uint) {
	t.Helper()

	if gotGameOver != wantGameOver {
		t.Fatalf("Game should not be over")
	}

	var diffs []uint
	for j, row := range gotBoard {
		for i, val := range row {
			if val != wantBoard[j][i] {
				diffs = append(diffs, val)
			}
		}
	}

	if newTile && len(diffs) != 1 {
		t.Fatalf("Invalid number of new values added after move. got %d, want 1", len(diffs))
	}

	for _, val := range diffs {
		if val != 2 && val != 4 {
			t.Fatalf("Invalid new value. got %d, want 2 or 4", val)
		}
	}

	if gotScore != wantScore {
		t.Fatalf("Invalid score. got %d, want %d", gotScore, wantScore)
	}
}
