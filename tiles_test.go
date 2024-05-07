package tiles

import (
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
		name  string
		board [][]uint
		down  moveRes
		up    moveRes
		left  moveRes
		right moveRes
	}{
		{
			name: "Sparse board",
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
			left: moveRes{
				want: [][]uint{
					{2, 4, 0, 0},
					{2, 4, 8, 0},
					{2, 4, 0, 0},
					{2, 4, 0, 0},
				},
				score:    0,
				gameOver: false,
				newTile:  true,
			},
			right: moveRes{
				want: [][]uint{
					{0, 0, 2, 4},
					{0, 2, 4, 8},
					{0, 0, 2, 4},
					{0, 0, 2, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  true,
			},
		},
		{
			name: "Filled board",
			board: [][]uint{
				{2, 4, 2, 4},
				{2, 4, 2, 4},
				{2, 4, 2, 4},
				{2, 4, 2, 4},
			},
			down: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{4, 8, 4, 8},
					{4, 8, 4, 8},
				},
				score:    48,
				gameOver: false,
				newTile:  true,
			},
			up: moveRes{
				want: [][]uint{
					{4, 8, 4, 8},
					{4, 8, 4, 8},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				score:    48,
				gameOver: false,
				newTile:  true,
			},
			left: moveRes{
				want: [][]uint{
					{2, 4, 2, 4},
					{2, 4, 2, 4},
					{2, 4, 2, 4},
					{2, 4, 2, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			right: moveRes{
				want: [][]uint{
					{2, 4, 2, 4},
					{2, 4, 2, 4},
					{2, 4, 2, 4},
					{2, 4, 2, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
		},
		{
			name: "Game over board",
			board: [][]uint{
				{64, 8, 16, 8},
				{64, 16, 8, 16},
				{32, 8, 16, 8},
				{8, 16, 8, 16},
			},
			down: moveRes{
				want: [][]uint{
					{0, 8, 16, 8},
					{128, 16, 8, 16},
					{32, 8, 16, 8},
					{8, 16, 8, 16},
				},
				score:    128,
				gameOver: true,
				newTile:  true,
			},
			up: moveRes{
				want: [][]uint{
					{128, 8, 16, 8},
					{32, 16, 8, 16},
					{8, 8, 16, 8},
					{0, 16, 8, 16},
				},
				score:    128,
				gameOver: false,
				newTile:  true,
			},
			left: moveRes{
				want: [][]uint{
					{64, 8, 16, 8},
					{64, 16, 8, 16},
					{32, 8, 16, 8},
					{8, 16, 8, 16},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			right: moveRes{
				want: [][]uint{
					{64, 8, 16, 8},
					{64, 16, 8, 16},
					{32, 8, 16, 8},
					{8, 16, 8, 16},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := &game{board: clone(test.board)}
			gameOver := g.MoveDown()
			assertMoveRes(t, g.board, test.down.want, gameOver, test.down.gameOver, test.down.newTile, g.score, test.down.score)

			g = &game{board: clone(test.board)}
			gameOver = g.MoveUp()
			assertMoveRes(t, g.board, test.up.want, gameOver, test.up.gameOver, test.up.newTile, g.score, test.up.score)

			g = &game{board: clone(test.board)}
			gameOver = g.MoveLeft()
			assertMoveRes(t, g.board, test.left.want, gameOver, test.left.gameOver, test.left.newTile, g.score, test.left.score)

			g = &game{board: clone(test.board)}
			gameOver = g.MoveRight()
			assertMoveRes(t, g.board, test.right.want, gameOver, test.right.gameOver, test.right.newTile, g.score, test.right.score)
		})
	}
}

func assertMoveRes(t *testing.T, gotBoard, wantBoard [][]uint, gotGameOver, wantGameOver, newTile bool, gotScore, wantScore uint) {
	t.Helper()

	if gotGameOver != wantGameOver {
		t.Fatalf("Invalid game over. got %t, want %t", gotGameOver, wantGameOver)
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

func clone(source [][]uint) [][]uint {
	dest := newBoard()
	for j, row := range source {
		for i, val := range row {
			dest[j][i] = val
		}
	}
	return dest
}
