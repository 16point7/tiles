package game

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	g := New()

	assertNewGame(t, g)
}

func TestReset(t *testing.T) {
	g := &game{
		board: [][]uint{
			{2, 2, 2, 2},
			{4, 4, 4, 4},
			{8, 8, 8, 8},
			{16, 16, 16, 16},
		},
		score: 0,
	}

	g.MoveLeft()

	var scoreWant uint = 120

	if g.Score() != scoreWant {
		t.Fatalf("Invalid score after move. got %d, want %d", g.Score(), scoreWant)
	}

	g.Reset()

	assertNewGame(t, g)
}

func assertNewGame(t *testing.T, g *game) {
	t.Helper()

	if g.GameOver() {
		t.Fatalf("New game must be in progress.")
	}

	if g.Score() != 0 {
		t.Fatalf("New game must have score of 0. got %d", g.Score())
	}

	if len(g.Board()) != 4 {
		t.Fatalf("Invalid number of rows. got %d, want %d", len(g.Board()), Side)
	}

	for j, row := range g.Board() {
		if len(row) != Side {
			t.Fatalf("Row %d has invalid number of columns. got %d, want %d", j, len(row), Side)
		}
	}

	var initVals []uint
	for _, row := range g.Board() {
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
			name: "Vertically filled board",
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
			name: "Horizontally filled board",
			board: [][]uint{
				{2, 2, 2, 2},
				{4, 4, 4, 4},
				{2, 2, 2, 2},
				{4, 4, 4, 4},
			},
			down: moveRes{
				want: [][]uint{
					{2, 2, 2, 2},
					{4, 4, 4, 4},
					{2, 2, 2, 2},
					{4, 4, 4, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			up: moveRes{
				want: [][]uint{
					{2, 2, 2, 2},
					{4, 4, 4, 4},
					{2, 2, 2, 2},
					{4, 4, 4, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			left: moveRes{
				want: [][]uint{
					{4, 4, 0, 0},
					{8, 8, 0, 0},
					{4, 4, 0, 0},
					{8, 8, 0, 0},
				},
				score:    48,
				gameOver: false,
				newTile:  true,
			},
			right: moveRes{
				want: [][]uint{
					{0, 0, 4, 4},
					{0, 0, 8, 8},
					{0, 0, 4, 4},
					{0, 0, 8, 8},
				},
				score:    48,
				gameOver: false,
				newTile:  true,
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
		{
			name: "No movement board",
			board: [][]uint{
				{0, 0, 0, 0},
				{2, 4, 8, 16},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			down: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{2, 4, 8, 16},
				},
				score:    0,
				gameOver: false,
				newTile:  true,
			},
			up: moveRes{
				want: [][]uint{
					{2, 4, 8, 16},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				score:    0,
				gameOver: false,
				newTile:  true,
			},
			left: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{2, 4, 8, 16},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			right: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{2, 4, 8, 16},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
		},
		{
			name: "No double collapse",
			board: [][]uint{
				{32, 16, 16, 8},
				{16, 4, 8, 16},
				{8, 2, 2, 4},
				{0, 0, 0, 0},
			},
			down: moveRes{
				want: [][]uint{
					{0, 0, 0, 0},
					{32, 16, 16, 8},
					{16, 4, 8, 16},
					{8, 2, 2, 4},
				},
				score:    0,
				gameOver: false,
				newTile:  true,
			},
			up: moveRes{
				want: [][]uint{
					{32, 16, 16, 8},
					{16, 4, 8, 16},
					{8, 2, 2, 4},
					{0, 0, 0, 0},
				},
				score:    0,
				gameOver: false,
				newTile:  false,
			},
			left: moveRes{
				want: [][]uint{
					{32, 32, 8, 0},
					{16, 4, 8, 16},
					{8, 4, 4, 0},
					{0, 0, 0, 0},
				},
				score:    36,
				gameOver: false,
				newTile:  true,
			},
			right: moveRes{
				want: [][]uint{
					{0, 32, 32, 8},
					{16, 4, 8, 16},
					{0, 8, 4, 4},
					{0, 0, 0, 0},
				},
				score:    36,
				gameOver: false,
				newTile:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := &game{board: clone(test.board)}
			g.MoveDown()
			assertMoveRes(t, g.Board(), test.down.want, g.GameOver(), test.down.gameOver, test.down.newTile, g.Score(), test.down.score)

			g = &game{board: clone(test.board)}
			g.MoveUp()
			assertMoveRes(t, g.Board(), test.up.want, g.GameOver(), test.up.gameOver, test.up.newTile, g.Score(), test.up.score)

			g = &game{board: clone(test.board)}
			g.MoveLeft()
			assertMoveRes(t, g.Board(), test.left.want, g.GameOver(), test.left.gameOver, test.left.newTile, g.Score(), test.left.score)

			g = &game{board: clone(test.board)}
			g.MoveRight()
			assertMoveRes(t, g.Board(), test.right.want, g.GameOver(), test.right.gameOver, test.right.newTile, g.Score(), test.right.score)
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
	} else if !newTile && len(diffs) != 0 {
		t.Fatalf("Invalid number of new values added after move. got %d, want 0", len(diffs))
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
	dest := createBoard()
	for j, row := range source {
		for i, val := range row {
			dest[j][i] = val
		}
	}
	return dest
}

func TestIsTerminal(t *testing.T) {
	tests := []struct {
		board      [][]uint
		isTerminal bool
	}{
		{
			board: [][]uint{
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			isTerminal: false,
		},
		{
			board: [][]uint{
				{2, 4, 8, 16},
				{16, 8, 4, 2},
				{2, 4, 8, 16},
				{16, 8, 4, 2},
			},
			isTerminal: true,
		},
		{
			board: [][]uint{
				{2, 4, 8, 16},
				{16, 0, 4, 2},
				{2, 4, 8, 16},
				{16, 8, 4, 2},
			},
			isTerminal: false,
		},
		{
			board: [][]uint{
				{0, 4, 8, 16},
				{16, 8, 4, 2},
				{2, 4, 8, 16},
				{16, 8, 4, 2},
			},
			isTerminal: false,
		},
		{
			board: [][]uint{
				{2, 4, 2, 4},
				{2, 4, 2, 4},
				{2, 4, 2, 4},
				{2, 4, 2, 4},
			},
			isTerminal: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got := isTerminal(test.board)
			if got != test.isTerminal {
				t.Fatalf("isTermina(board) = %t, want %t", got, test.isTerminal)
			}
		})
	}
}
